package controller

import (
	"log"
	"GP/constant"
	"io/ioutil"
	"net/http"
	"GP/utils"
	"encoding/json"
	"net/url"
	"GP/services/comment"
)

func GetCommentList(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	queryForm,err := url.ParseQuery(r.URL.RawQuery)
	username := queryForm["username"][0]

	if len(username) <= 0 {
		errmsg := "username is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	commentinfo, err := comment.GetCommentList(username)
	if err != nil {
		errmsg := "commentinfo from database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("get commentlist success").FbData(commentinfo).Response()
}

type NewCommentParams struct {
	UserName     string `json:"username"`
	FromUserName string `json:"fromusername"`
	FromNickName string `json:"fromnickname"`
	Comment      string `json:"comment"`
	Time         string `json:"time"`
}

func NewComment(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errmsg := "read body error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	params := &NewCommentParams{}
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

	if len(params.FromUserName) <= 0 {
		errmsg := "FromUserName is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	if len(params.FromNickName) <= 0 {
		errmsg := "FromNickName is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	if len(params.Comment) <= 0 {
		errmsg := "Comment is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	if len(params.Time) <= 0 {
		errmsg := "Time is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	err = comment.NewComment(params.UserName, params.FromUserName, params.FromNickName, params.Comment, params.Time)
	if err != nil {
		errmsg := "NewComment into database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("new comment success").Response()
}