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
	params := RoomNameParams{}
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
	fb.FbCode(constant.SUCCESS).FbMsg("create room success").Response()
}

func GetOneRoom(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	queryForm,err := url.ParseQuery(r.URL.RawQuery)
	roomname := queryForm["roomname"][0]
	oneroominfo, err := room.GetOneRoom(roomname)
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

func CancelBanRoom(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	params := &RoomNameParams {}
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

	err = room.CancelBanRoom(params.RoomName)
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
	params := &RoomNameParams {}
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

	err = room.BanRoom(params.RoomName)
	if err != nil {
		errmsg := "Banroom into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("banroom success").Response()
}