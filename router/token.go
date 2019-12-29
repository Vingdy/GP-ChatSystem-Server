package router

import (
	"GP/constant"
	"GP/model"
	"GP/utils"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"net/http"
)

func TokenCheck(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fb := utils.NewFeedBack(w)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errmsg := "read body error:" + err.Error()
			log.Println(errmsg)
			fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
			return
		}
		check := &model.Token{}
		err = json.Unmarshal(body, check)
		if err != nil {
			errmsg := "json unmarshal error:" + err.Error()
			log.Println(errmsg)
			fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
			return
		}
		authorization := check.AccessToken
		token, err := jwt.Parse(authorization, func(token *jwt.Token) (i interface{}, e error) {
			return []byte("vector.sign"), nil
		})
		if err != nil {
			if err, ok := err.(*jwt.ValidationError); ok {
				if err.Errors&jwt.ValidationErrorMalformed != 0 {
					return
				}
				if err.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					fmt.Println(err)
					return
				}
			}
			return
		}
		finToken := token.Claims.(jwt.MapClaims) // 获取token里面的字段，如生成填入的username
		fmt.Println(finToken)
	})
}

