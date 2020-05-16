package controller

import (
	"net/http"
	"GP/utils"
	"io/ioutil"
	"log"
	"GP/constant"
	"encoding/json"
	"GP/services/user"
	"net/url"
)

type UserIdParams struct {
	UserId string `json:"id"`
}

type UserNameParams struct {
	UserName string `json:"username"`
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if queryForm["id"] == nil {
		errmsg := "query id is nil"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}
	id := queryForm["id"][0]

	if len(id) <= 0 {
		errmsg := "id is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	userinfo, err := user.GetOneUser(id)
	if err != nil {
		errmsg := "GetOneUser from database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	if len(userinfo) == 0 {
		fb.FbCode(constant.PASSWORD_NOT_RIGHT).FbMsg("GetOneUser id not right").Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("get one user success").FbData(userinfo[0]).Response()
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)

	userinfo, err := user.GetUserList()
	if err != nil {
		errmsg := "GetUserList from database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("get user list success").FbData(userinfo).Response()
}

type UpdateUserParams struct {
	UserId string `json:"id"`
	NickName string `json:"nickname"`
	Phone 	 string `json:"phone"`
	Label 	 string `json:"label"`
	FontType 	 string `json:"fonttype"`
	FontColor string `json:"fontcolor"`
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	params := &UpdateUserParams{}
	err = json.Unmarshal(body, params)
	if err != nil {
		errmsg := "json unmarshal error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(params.UserId) <= 0 {
		errmsg := "UserId is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	if len(params.NickName) <= 0 {
		errmsg := "NickName is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	if len(params.Phone) <= 0 {
		errmsg := "Phone is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	if len(params.Label) <= 0 {
		errmsg := "Label is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	if len(params.FontType) <= 0 {
		errmsg := "FontType is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	if len(params.FontColor) <= 0 {
		errmsg := "FontColor is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	err = user.UpdateUser(params.UserId, params.NickName, params.Phone, params.Label, params.FontType, params.FontColor)
	if err != nil {
		errmsg := "UpdateUser into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("update user success").Response()
}

type UpdatePasswordParams struct {
	UserId string `json:"id"`
	PassWord string `json:"password"`
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	params := &UpdatePasswordParams{}
	err = json.Unmarshal(body, params)
	if err != nil {
		errmsg := "json unmarshal error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(params.UserId) <= 0 {
		errmsg := "UserId is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	if len(params.PassWord) <= 0 {
		errmsg := "Password is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	err = user.UpdatePassword(params.UserId, params.PassWord)
	if err != nil {
		errmsg := "UpdatePassword into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("update password success").Response()
}

func BanUser(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	params := &UserIdParams {}
	err = json.Unmarshal(body, params)
	if err != nil {
		errmsg := "json unmarshal error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(params.UserId) <= 0 {
		errmsg := "UserId is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	err = user.BanUser(params.UserId)
	if err != nil {
		errmsg := "Banuser into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("banuser success").Response()
}

func CancelBanUser(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	params := &UserIdParams {}
	err = json.Unmarshal(body, params)
	if err != nil {
		errmsg := "json unmarshal error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(params.UserId) <= 0 {
		errmsg := "UserId is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	err = user.CancelBanUser(params.UserId)
	if err != nil {
		errmsg := "CancelBanUser into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("cancel banuser success").Response()
}

func UpUserRole(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	params := &UserIdParams{}
	err = json.Unmarshal(body, params)
	if err != nil {
		errmsg := "json unmarshal error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(params.UserId) <= 0 {
		errmsg := "UserId is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	err = user.UpUserRole(params.UserId)
	if err != nil {
		errmsg := "UpUserRole into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("upuserrole success").Response()
}

func DownUserRole(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	params := &UserIdParams{}
	err = json.Unmarshal(body, params)
	if err != nil {
		errmsg := "json unmarshal error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(params.UserId) <= 0 {
		errmsg := "UserId is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	err = user.DownUserRole(params.UserId)
	if err != nil {
		errmsg := "DownUserRole into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("downuserrole success").Response()
}

type FindStringParams struct {
	FindString string `json:"findstring"`
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	queryForm,err := url.ParseQuery(r.URL.RawQuery)
	findstring := queryForm["findstring"][0]
	userinfo, err := user.FindUser(findstring)
	if err != nil {
		errmsg := "FindUser into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("finduser success").FbData(userinfo).Response()
}

func GetUserRole(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	queryForm,err := url.ParseQuery(r.URL.RawQuery)
	username := queryForm["username"][0]
	role, err := user.GetUserRole(username)
	if err != nil {
		errmsg := "getuserrole from database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("get user role success").FbData(role).Response()
}

