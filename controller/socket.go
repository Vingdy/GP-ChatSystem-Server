package controller

import (
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
	"encoding/json"
)

type ConnInfo struct {
	Token string
	Name string
	Conn *websocket.Conn
}

type Message struct {
	Type    string `json:type`      //消息类型
	Name    string `json:"name"`    // 用户名称
	Message string `json:"message"` // 消息内容
}

var ClientConnsMap map[int]ConnInfo

var joinchan chan ConnInfo
var leavechan chan ConnInfo
var messagechan chan Message

type Socket struct {

}

func init() {
	ClientConnsMap = make(map[int]ConnInfo)
	joinchan = make(chan ConnInfo, 10)
	leavechan = make(chan ConnInfo, 10)
	messagechan = make(chan Message, 50)
	go MessageHandle()
}

func MessageHandle() {
	for {
		select {
			case msg := <- messagechan: {
				for _, client := range ClientConnsMap {
					data, err := json.Marshal(msg)
					if err != nil {
						return
					}
					if client.Conn.WriteMessage(websocket.TextMessage, data) != nil {
						fmt.Errorf("fail to write message")
					}
				}
			}
			case client := <- joinchan: {
				ClientConnsMap[1] = client
				var msg Message
				msg.Type = "1"
				msg.Name = client.Name
				msg.Message = fmt.Sprintf("%s加入了房间", client.Name)
				messagechan <- msg
			}
			case client := <- leavechan: {
				if _, find := ClientConnsMap[1]; !find {
					break
				}
				delete(ClientConnsMap, 1)
				var msg Message
				msg.Name = client.Name
				msg.Type = "2"
				msg.Message = fmt.Sprintf("%s离开了房间", client.Name)
				messagechan <- msg
			}
		}
	}
}

func (s Socket) NewSocket(token string, w http.ResponseWriter, r *http.Request) (client *ConnInfo) {
	ws := websocket.Upgrader{
		ReadBufferSize:4096,
		WriteBufferSize:4096,
		CheckOrigin:func(r *http.Request) bool {
			return true
		},
	}

	conn, err := ws.Upgrade(w, r, w.Header())
	if err != nil {
		return
	}
	//fmt.Println(conn)
	client = &ConnInfo{
		Token:token,
		Conn:conn,
	}
	return client
}

func WsMain(w http.ResponseWriter, r *http.Request) {
	if !websocket.IsWebSocketUpgrade(r) {
		return
	}

	accessToken := r.Header.Get("AccessToken")

	s := new(Socket)
	newclient := s.NewSocket(accessToken, w, r)
	//fmt.Println(newclient)
	if _, find := ClientConnsMap[1]; !find {
		joinchan <- *newclient
		fmt.Println("connet success")
	}
	defer func() {
		leavechan <- *newclient
		newclient.Conn.Close()
	}()
	//对于这个goroutinue保持监听
	for {
		_, msgstr, err := newclient.Conn.ReadMessage()
		if err != nil {
			break
		}
		var msg Message
		msg.Type = "0"
		msg.Name = newclient.Name
		msg.Message = string(msgstr)
		messagechan <- msg
	}
}

