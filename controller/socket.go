package controller

import (
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
	"encoding/json"
	"strconv"
	"time"
	"GP/services/room"
	"log"
	"GP/utils"
	"GP/services/history"
)

type ConnInfo struct {
	Id         int
	Token      string
	Name       string
	Conn       *websocket.Conn
	Room       string
	SendChan   *chan Message
}

type Message struct {
	Type      string `json:"type"`               //消息类型
	EndPoint  string `json:"endpoint,omitempty"` //发送目标
	Name      string `json:"name"`               // 用户名称
	RealName      string `json:"realname"`
	RoomName  string `json:"roomname"`
	Message   string `json:"message"` // 消息内容
	Label     string `json:"label"`
	FontType  string `json:"fonttype"`
	FontColor string `json:"fontcolor"`
	NowMember []string `json:"nowmember"`
	Time      string `json:"time"` //发送时间
	Token     string `json:token`
}

type Room struct {
	Name           string
	ClientConnsMap map[int]ConnInfo
	Joinchan       chan ConnInfo
	Leavechan      chan ConnInfo
	Messagechan    chan Message
}

var AllClient map[int]*ConnInfo
var AllJoinchan chan ConnInfo
var AllLeavechan chan ConnInfo
var AllMessagechan chan Message

var Rooms []Room

type Socket struct {

}

func Ws_init() {
	roomlist, err := room.GetRoomList()
	if err != nil {
		errmsg := "GetRoomList from database error:" + err.Error()
		log.Println(errmsg)
		return
	}
	for i := 0; i < len(roomlist); i++ {
		newRoom := Room{
			Name:           roomlist[i].RoomName,
			ClientConnsMap: make(map[int]ConnInfo),
			Joinchan:       make(chan ConnInfo, 20),
			Leavechan:      make(chan ConnInfo, 20),
			Messagechan:    make(chan Message, 60),
		}
		Rooms = append(Rooms, newRoom)
		go newRoom.MessageHandle()
	}
	AllClient = make(map[int]*ConnInfo)
	AllJoinchan = make(chan ConnInfo, 20)
	AllLeavechan = make(chan ConnInfo, 20)
	go MessageHandle()
}

func MessageHandle() {
	for {
		select {
		case client := <-AllJoinchan:
			{
				AllClient[client.Id] = &client
			}
		case client := <-AllLeavechan:
			{
				if _, find := AllClient[client.Id]; !find {
					break
				}
				delete(AllClient, client.Id)
			}
		}
	}
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
					} else {
						if msg.Type == "message" {
							err = history.NewHistory(r.Name, msg.Name, msg.Message, msg.Label, msg.FontType, msg.FontColor,msg.Time)
							log.Println(err)
						}
					}
				}
			}
			case client := <- r.Joinchan: {
				r.ClientConnsMap[client.Id] = client
				var msg Message
				msg.Type = "join"
				msg.RoomName = client.Room
				msg.Name = client.Name
				var newmember []string
				for _, v := range r.ClientConnsMap {
					newmember = append(newmember, v.Name)
				}
				msg.NowMember = newmember
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

func (s Socket) NewSocket(token string, userid string, username string, roomname string, sendchan *chan Message,  w http.ResponseWriter, r *http.Request) (client ConnInfo, err error) {
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
		return client, err
	}
	id, _ := strconv.Atoi(userid)
	client = ConnInfo{
		Token:token,
		Id: id,
		Name:username,
		Conn:conn,
		Room:roomname,
		SendChan:sendchan,
	}
	return client, nil
}

var count int

func WsMain(w http.ResponseWriter, r *http.Request) {
	if !websocket.IsWebSocketUpgrade(r) {
		return
	}

	accessToken := r.Header.Get("Sec-WebSocket-Protocol")
	fmt.Println("accessToken", accessToken)
	userinfo, _ := utils.GetTokenInfo(accessToken)
	fmt.Println(userinfo)
	//测试用数据
	/*userinfo := model.User{
		Id:strconv.Itoa(count),
		NickName:strconv.Itoa(count),
	}
	count++*/


	s := new(Socket)

	var client ConnInfo
	var nowroom Room
	fmt.Println(Rooms)
	for _, room := range Rooms {
		if room.Name == "公共房间" {
			newclient, err := s.NewSocket(accessToken, userinfo.Id, userinfo.NickName, room.Name, &room.Messagechan, w, r)
			if err != nil {
				fmt.Println("newclient create fail")
				return
			}
			client = newclient
			nowroom = room
		}
	}
	if client.Conn == nil {
		fmt.Println("client conn is nil")
		return
	}
	fmt.Println(client)
	id, _ := strconv.Atoi(userinfo.Id)
	if _, find := nowroom.ClientConnsMap[id]; !find {
		nowroom.Joinchan <- client
		AllJoinchan <- client
		fmt.Println("connet success")
	}
	defer func() {
		nowroom.Leavechan <- client
		AllLeavechan <- client
		client.Conn.Close()
	}()

	//fmt.Println(newclient)

	//对于这个goroutinue保持监听
	for {
		_, data, err := client.Conn.ReadMessage()
		accessToken := r.Header.Get("Sec-WebSocket-Protocol")
		fmt.Println("accessToken222", accessToken)
		userinfo, _ := utils.GetTokenInfo(accessToken)
		fmt.Println(string(data), nowroom)
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
		if msg.Type == "change" {
			fmt.Println("change", msg)
			nowroom.Leavechan <- client
			for _, room := range Rooms {
				if room.Name == msg.Message {
					nowroom = room
					client.SendChan = &room.Messagechan
					client.Room = room.Name
					fmt.Println(client.Room)
					room.Joinchan <- client
				}
			}
		} else {
			msg.RoomName = nowroom.Name
			msg.Label = userinfo.Label
			msg.RealName= userinfo.UserName
			msg.FontType = userinfo.FontType
			msg.FontColor = userinfo.FontColor
			var newmember []string
			for _, v := range nowroom.ClientConnsMap {
				newmember = append(newmember, v.Name)
			}
			msg.NowMember = newmember
			msg.Token = accessToken
			nowroom.Messagechan <- msg
		}
	}
}

