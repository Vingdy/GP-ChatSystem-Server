package utils

import (
	"GP/model"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type FeedBack model.FeedBack

func NewFeedBack(w http.ResponseWriter) *FeedBack {
	return &FeedBack{Dist:w}
}

func (f *FeedBack) FbDist(w http.ResponseWriter) *FeedBack {
	f.Dist = w
	return f
}

func (f *FeedBack) FbCode(code int) *FeedBack {
	f.Code = code
	return f
}

func (f *FeedBack) FbMsg(msg string) *FeedBack {
	f.Msg = msg
	return f
}

func (f *FeedBack) FbData(data interface{}) *FeedBack {
	f.Data = data
	return f
}

func (f *FeedBack)Response()(err error) {
	//out:=&FeedBack
	if f.Dist == nil {
		return errors.New("DistWriter is empty")
	}
	result, err := json.Marshal(f)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(f.Dist, string(result))
	f.Clear()
	return nil
}

func (f *FeedBack) Clear() {
	f.Data = nil
	f.Msg = ""
	f.Code = 0
}

