package controller

import (
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
	"encoding/json"
	"strconv"
	"GP/model"
	"time"
)

type ConnInfo struct {
	Id       int
	Token    string
	Name     string
	Conn     *websocket.Conn
	Room     string
	SendChan *chan Message
}

type Message struct {
	Type     string `json:"type"`               //消息类型
	EndPoint string `json:"endpoint,omitempty"` //发送目标
	Name     string `json:"name"`               // 用户名称
	Message  string `json:"message"`            // 消息内容
	Time     string `json:"time"`               //发送时间
}

type Room struct {
	Name           string
	ClientConnsMap map[int]ConnInfo
	Joinchan       chan ConnInfo
	Leavechan      chan ConnInfo
	Messagechan    chan Message
}

/*var Joinchan chan ConnInfo
var Leavechan chan ConnInfo
var Messagechan chan Message*/

var Rooms []Room

type Socket struct {

}

func init() {
	newRoom := Room{
		Name:           "世界",
		ClientConnsMap: make(map[int]ConnInfo),
		Joinchan:       make(chan ConnInfo, 10),
		Leavechan:      make(chan ConnInfo, 10),
		Messagechan:    make(chan Message, 50),
	}
	Rooms = append(Rooms, newRoom)
	/*Joinchan = make(chan ConnInfo, 10)
	Leavechan = make(chan ConnInfo, 10)
	Messagechan = make(chan Message, 50)*/
	go newRoom.MessageHandle()
}

func (r Room)MessageHandle() {
	for {
		select {
			case msg := <- r.Messagechan: {
				for _, client := range r.ClientConnsMap {
					msg.Time = time.Now().Format("2006-01-02 15:04:05")
					data, err := json.Marshal(msg)
					if err != nil {
						return
					}
					fmt.Println(client, string(data))
					if client.Conn.WriteMessage(websocket.TextMessage, data) != nil {
						fmt.Errorf("fail to write message")
					}
				}
			}
			case client := <- r.Joinchan: {
				r.ClientConnsMap[client.Id] = client
				var msg Message
				msg.Type = "join"
				msg.Name = client.Name
				msg.Message = fmt.Sprintf("%s加入了房间", client.Name)
				r.Messagechan <- msg
			}
			case client := <- r.Leavechan: {
				if _, find := r.ClientConnsMap[client.Id]; !find {
					break
				}
				delete(r.ClientConnsMap, client.Id)
				var msg Message
				msg.Name = client.Name
				msg.Type = "leave"
				msg.Message = fmt.Sprintf("%s离开了房间", client.Name)
				r.Messagechan <- msg
			}
		}
	}
}

func (s Socket) NewSocket(token string, userid string, username string, roomname string, sendchan *chan Message,  w http.ResponseWriter, r *http.Request) (client *ConnInfo) {
	ws := websocket.Upgrader{
		ReadBufferSize:4096,
		WriteBufferSize:4096,
		CheckOrigin:func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{r.Header.Get("Sec-WebSocket-Protocol")},
	}

	conn, err := ws.Upgrade(w, r, w.Header())
	if err != nil {
		return
	}
	id, _ := strconv.Atoi(userid)
	//fmt.Println(conn)
	client = &ConnInfo{
		Token:token,
		Id: id,
		Name:username,
		Conn:conn,
		Room:roomname,
		SendChan:sendchan,
	}
	return client
}

var count int

func WsMain(w http.ResponseWriter, r *http.Request) {
	if !websocket.IsWebSocketUpgrade(r) {
		return
	}

	accessToken := r.Header.Get("AccessToken")
	//userinfo, _ := utils.GetTokenInfo(accessToken)

	//测试用数据
		userinfo := model.User{
			Id:strconv.Itoa(count),
			NickName:strconv.Itoa(count),
		}
		count++


	s := new(Socket)

	var newclient *ConnInfo
	var nowroom Room
	for _, room := range Rooms {
		if room.Name == "世界" {
			newclient = s.NewSocket(accessToken, userinfo.Id, userinfo.NickName, room.Name, &room.Messagechan, w, r)
			nowroom = room
		}
	}
	if newclient.Conn == nil {
		fmt.Println("client conn is nil")
		return
	}

	id, _ := strconv.Atoi(userinfo.Id)
	if _, find := nowroom.ClientConnsMap[id]; !find {
		nowroom.Joinchan <- *newclient
		fmt.Println("connet success")
	}
	defer func() {
		nowroom.Leavechan <- *newclient
		newclient.Conn.Close()
	}()

	//fmt.Println(newclient)

	//对于这个goroutinue保持监听
	for {
		_, data, err := newclient.Conn.ReadMessage()
		fmt.Println(string(data))
		if err != nil {
			fmt.Println(err)
			break
		}
		var msg Message
		err = json.Unmarshal(data, &msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		nowroom.Messagechan <- msg
	}
}

