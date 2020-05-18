package controller

import (
	"GP/constant"
	"GP/services/friend"
	"GP/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GetCheckFriend(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	username := queryForm["username"][0]

	if len(username) <= 0 {
		errmsg := "username is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	friendinfo, err := friend.GetCheckFriend(username)
	if err != nil {
		errmsg := "friendinfo from database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("get checkfriend success").FbData(friendinfo).Response()
}

func GetFriendList(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	username := queryForm["username"][0]

	if len(username) <= 0 {
		errmsg := "username is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	friendinfo, err := friend.GetFriendList(username)
	if err != nil {
		errmsg := "friendinfo from database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("get friendlist success").FbData(friendinfo).Response()
}

type NewFriendParams struct {
	UserName1 string `json:"username1"`
	NickName1 string `json:"nickname1"`
	Id2       string `json:"id2"`
	UserName2 string `json:"username2"`
	NickName2 string `json:"nickname2"`
	Label2    string `json:"label2"`
}

func NewFriend(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	params := &NewFriendParams{}
	err = json.Unmarshal(body, params)
	if err != nil {
		errmsg := "json unmarshal error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(params.UserName1) <= 0 {
		errmsg := "UserName1 is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	if len(params.NickName1) <= 0 {
		errmsg := "NickName1 is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	if len(params.UserName2) <= 0 {
		errmsg := "UserName2 is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	if len(params.NickName2) <= 0 {
		errmsg := "NickName2 is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	if len(params.Label2) <= 0 {
		errmsg := "Label2 is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	find, err := friend.NewFriendCheck(params.UserName1, params.UserName2)
	if find {
		fb.FbCode(constant.NEWFRIEND_HAS_BEEN_EXIST).FbMsg("newfriend has been exist").Response()
		return
	}
	if err != nil {
		errmsg := "NewFriendCheck data write into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	id, err := friend.NewFriend(params.UserName1, params.NickName1, params.Id2, params.UserName2, params.NickName2, params.Label2)
	if err != nil {
		errmsg := "NewFriend into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("new friend success").FbData(id).Response()
}

type PassFriendParams struct {
	Id string `json:"id"`
}

func PassFriend(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	params := &PassFriendParams{}
	err = json.Unmarshal(body, params)
	if err != nil {
		errmsg := "json unmarshal error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(params.Id) < 0 {
		errmsg := "Id is not right"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	find, err := friend.PassFriendIdCheck(params.Id)
	if !find {
		fb.FbCode(constant.PASSFRIEND_NOT_EXIST).FbMsg("id not exist").Response()
		return
	}
	if err != nil {
		errmsg := "PassFriendIdCheck data write into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	err = friend.PassFriend(params.Id)
	if err != nil {
		errmsg := "PassFriend into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("pass friend success").Response()
}

func UnPassFriend(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	params := &PassFriendParams{}
	err = json.Unmarshal(body, params)
	if err != nil {
		errmsg := "json unmarshal error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(params.Id) < 0 {
		errmsg := "Id is not right"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	find, err := friend.PassFriendIdCheck(params.Id)
	if !find {
		fb.FbCode(constant.PASSFRIEND_NOT_EXIST).FbMsg("id not exist").Response()
		return
	}
	if err != nil {
		errmsg := "PassFriendIdCheck data write into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	err = friend.UnPassFriend(params.Id)
	if err != nil {
		errmsg := "UnPassFriend into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("unpass friend success").Response()
}
