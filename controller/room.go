package controller

import (
	"log"
	"GP/constant"
	"net/http"
	"GP/utils"
	"io/ioutil"
	"encoding/json"
	"net/url"
	"GP/services/room"
)

type RoomIdParams struct {
	RoomId string `json:"id"`
}

type RoomNameParams struct {
	RoomName string `json:"roomname"`
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	params := &RoomNameParams{}
	err = json.Unmarshal(body, params)
	if err != nil {
		errmsg := "json unmarshal error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(params.RoomName) <= 0 {
		errmsg := "RoomName is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	err = room.CreateRoom(params.RoomName)
	if err != nil {
		errmsg := "CreateRoom into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	newRoom := Room{
		Name:           params.RoomName,
		ClientConnsMap: make(map[int]ConnInfo),
		Joinchan:       make(chan ConnInfo, 20),
		Leavechan:      make(chan ConnInfo, 20),
		Messagechan:    make(chan Message, 60),
	}
	Rooms = append(Rooms, newRoom)
	go newRoom.MessageHandle()
	fb.FbCode(constant.SUCCESS).FbMsg("create room success").Response()
}

func GetOneRoom(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	queryForm,err := url.ParseQuery(r.URL.RawQuery)
	id := queryForm["id"][0]
	oneroominfo, err := room.GetOneRoom(id)
	if err!=nil {
		errmsg := "GetOneRoom from database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("getoneroom success").FbData(oneroominfo[0]).Response()
}

func GetRoomList(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	roomlist, err := room.GetRoomList()
	if err!=nil {
		errmsg := "GetRoomList from database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("get room list success").FbData(roomlist).Response()
}

func GetUseRoomList(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	roomlist, err := room.GetUseRoomList()
	if err!=nil {
		errmsg := "GetUserRoomList from database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("get use room list success").FbData(roomlist).Response()
}

func CancelBanRoom(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	params := &RoomIdParams {}
	err = json.Unmarshal(body, params)
	if err != nil {
		errmsg := "json unmarshal error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(params.RoomId) <= 0 {
		errmsg := "RoomId is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	err = room.CancelBanRoom(params.RoomId)
	if err != nil {
		errmsg := "Banroom into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("cancel banroom success").Response()
}

func BanRoom(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	params := &RoomIdParams {}
	err = json.Unmarshal(body, params)
	if err != nil {
		errmsg := "json unmarshal error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(params.RoomId) <= 0 {
		errmsg := "RoomId is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	err = room.BanRoom(params.RoomId)
	if err != nil {
		errmsg := "Banroom into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("banroom success").Response()
}