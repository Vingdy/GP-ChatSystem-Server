package controller

import (
	"GP/constant"
	"GP/model"
	"GP/services/register"
	"GP/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:"+err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	newUser := &model.Register{}
	err = json.Unmarshal(body, newUser)
	if err != nil {
		errmsg := "json unmarshal error:"+err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(newUser.UserName)<=0{
		errmsg := "UserName is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}
	if len(newUser.PassWord)<=0{
		errmsg := "PassWord is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}
	if len(newUser.NickName)<=0{
		errmsg := "NickName is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	find, err := register.RegisterAccCheck(newUser.UserName)
	if find {
		fb.FbCode(constant.ACCOUNT_HAS_BEEN_REGISTER).FbMsg("account has been register").Response()
		return
	}
	if err != nil {
		errmsg := "RegisterAccCheck data write into database error:"+err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	err = register.Register(newUser.UserName, newUser.PassWord, newUser.NickName)
	if err != nil {
		errmsg := "Register data write into database error:"+err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("register success").Response()
}