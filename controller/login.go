package controller

import (
	"GP/constant"
	"GP/model"
	"GP/services/login"
	"GP/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"GP/redis"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	loginUser := &model.Login{}
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
	if len(loginUser.PassWord) <= 0 {
		errmsg := "PassWord is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	find, err := login.LoginAccCheck(loginUser.UserName)
	if !find {
		fb.FbCode(constant.ACCOUNT_HAS_BEEN_EXIST).FbMsg("account has been exist").Response()
		return
	}
	if err != nil {
		errmsg := "LoginAccCheck data write into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	userinfo, err := login.Login(loginUser.UserName, loginUser.PassWord)
	if len(userinfo) == 0 {
		fb.FbCode(constant.PASSWORD_NOT_RIGHT).FbMsg("password not right").Response()
		return
	}
	if err != nil {
		errmsg := "Login data write into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}

	//生成一条随机数据
	/*signKey := sid.Id()//sid:8 byte time - 8 byte random
	signKey = signKey[:32]
	hashData, err := utils.AesEncryptCBC([]byte(userinfo[0].Id), []byte(signKey))
	if err != nil {
		errmsg := "hashData create fail:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	Data, err := utils.AesDecryptCBC([]byte(hashData), []byte(signKey))
	if err != nil {
		errmsg := "hashData create fail:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fmt.Println(string(hashData),string(Data),signKey)
	u,_ := uuid.NewV4()
	fmt.Println(u)*/
	newToken := model.Token{}
	//重新给token时间30min
	expired := time.Now().Add(60 * time.Minute).Unix()
	accessClaims := &jwt.StandardClaims{
		ExpiresAt: expired,            // 过期时间，必须设置
		Issuer:    loginUser.UserName, // 可不必设置，也可以填充用户名，
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)   //生成token
	accessToken, err := token.SignedString([]byte("asign")) //签名
	if err != nil {
		errmsg := "token create error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	//重新登录时间15天
	/*refreshed := time.Now().Add(30 * time.Minute).Unix()
	refreshClaims := &jwt.StandardClaims{
		ExpiresAt: refreshed,            // 过期时间，必须设置
		Issuer:    loginUser.UserName, // 可不必设置，也可以填充用户名，
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)   //生成token
	refreshToken, err := token.SignedString([]byte("asign")) //签名
	if err != nil {
		return
	}*/

	/*newToken.RefreshToken = refreshToken
	newToken.RefreshAt = refreshed*/
	newToken.UserName = userinfo[0].UserName
	newToken.ExpiresAt = expired
	newToken.AccessToken = accessToken
	newToken.CreateAt = time.Now().Unix()
	//fb.FbCode(constant.SUCCESS).FbMsg("login success").Response()
	buffer, err := json.Marshal(userinfo[0])
	if err != nil {
		errmsg := "buffer json marshal failed:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	if err = redis.Redis.Set(accessToken, buffer, 60 * time.Minute).Err();err != nil {
		errmsg := "redis set token failed:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	w.Header().Set("AccessToken", accessToken)
	userinfo[0].Token = accessToken
	fb.FbCode(constant.SUCCESS).FbMsg("login success").FbData(userinfo[0]).Response()
}
