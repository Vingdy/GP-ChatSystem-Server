package controller

import (
	"GP/constant"
	"GP/model"
	"GP/services"
	"GP/utils"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:"+err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	loginUser := &model.Login{}
	err = json.Unmarshal(body, loginUser)
	if err != nil {
		errmsg := "json unmarshal error:"+err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	if len(loginUser.UserName)<=0{
		errmsg := "UserName is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}
	if len(loginUser.PassWord)<=0{
		errmsg := "PassWord is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	find, err := services.LoginAccCheck(loginUser.UserName)
	if !find {
		fb.FbCode(constant.ACCOUNT_HAS_BEEN_EXIST).FbMsg("account has been exist").Response()
		return
	}
	if err != nil {
		errmsg := "LoginAccCheck data write into database error:"+err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	userinfo,err := services.Login(loginUser.UserName, loginUser.PassWord)
	if len(userinfo) == 0 {
		fb.FbCode(constant.PASSWORD_NOT_RIGHT).FbMsg("password not right").Response()
		return
	}
	if err != nil {
		errmsg := "Login data write into database error:"+err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	newToken := model.Token{}
	claims := &jwt.StandardClaims{
		ExpiresAt:time.Now().Add(30*time.Second).Unix(), // 过期时间，必须设置
		Issuer:userinfo[0].UserName,   // 可不必设置，也可以填充用户名，
	}
	expired := time.Now().Add(148 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims) //生成token
	accessToken, err := token.SignedString([]byte("vector.sign"))//签名
	if err != nil {
		return
	}
	newToken.ExpiresAt = expired
	newToken.AccessToken = accessToken
	newToken.Timestamp = time.Now().Unix()

	fb.FbCode(constant.SUCCESS).FbMsg("login success").FbData(newToken).Response()
}
