package controller

import (
	"net/http"
	"GP/utils"
	"io/ioutil"
	"log"
	"GP/constant"
	"encoding/json"
	"GP/services/user"
	"GP/model"
)

type UserNameParams struct {
	UserName string `json:"username"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	loginUser := &UserNameParams{}
	err = json.Unmarshal(body, loginUser)
	if err != nil {
		errmsg := "json unmarshal error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(loginUser.UserName) <= 0 {
		errmsg := "UserName is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	userinfo, err := user.GetUser(loginUser.UserName)
	if len(userinfo) == 0 {
		fb.FbCode(constant.PASSWORD_NOT_RIGHT).FbMsg("GetUser username not right").Response()
		return
	}
	if err != nil {
		errmsg := "GetUser from database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("get user success").FbData(userinfo[0]).Response()
}

type UpdateUserParams struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Phone 	 string `json:"phone"`
	Label 	 string `json:"label"`
	Head 	 string `json:"head"`
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

	if len(params.UserName) <= 0 {
		errmsg := "UserName is empty"
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

	if len(params.Head) <= 0 {
		errmsg := "Head is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	err = user.UpdateUser(params.UserName, params.NickName, params.Phone, params.Label, params.Label)
	if err != nil {
		errmsg := "UpdateUser into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("update user success").Response()
}

type UpdatePasswordParams struct {
	UserName string `json:"username"`
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
	params := &model.Login{}
	err = json.Unmarshal(body, params)
	if err != nil {
		errmsg := "json unmarshal error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(params.UserName) <= 0 {
		errmsg := "UserName is empty"
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

	err = user.UpdatePassword(params.UserName, params.PassWord)
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
	params := &UserNameParams {}
	err = json.Unmarshal(body, params)
	if err != nil {
		errmsg := "json unmarshal error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(params.UserName) <= 0 {
		errmsg := "UserName is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	err = user.BanUser(params.UserName)
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
	params := &UserNameParams {}
	err = json.Unmarshal(body, params)
	if err != nil {
		errmsg := "json unmarshal error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(params.UserName) <= 0 {
		errmsg := "UserName is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	err = user.CancelBanUser(params.UserName)
	if err != nil {
		errmsg := "Banuser into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("cancel banuser success").Response()
}


