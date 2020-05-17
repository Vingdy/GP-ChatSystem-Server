package controller

import (
	"log"
	"GP/constant"
	"net/http"
	"GP/utils"
	"net/url"
	"GP/services/history"
)

func GetHistoryList(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	queryForm,err := url.ParseQuery(r.URL.RawQuery)
	roomname := queryForm["roomname"][0]

	if len(roomname) <= 0 {
		errmsg := "roomname is empty"
		log.Println(errmsg)
		fb.FbCode(constant.PARMAS_EMPTY).FbMsg(errmsg).Response()
		return
	}

	historyinfo, err := history.GetHistoryList(roomname)
	if err != nil {
		errmsg := "historyinfo from database error:" + err.Error()
		log.Println(errmsg)
		fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbMsg(errmsg).Response()
		return
	}
	fb.FbCode(constant.SUCCESS).FbMsg("get historylist success").FbData(historyinfo).Response()
}