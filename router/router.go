package router

import (
	"GP/controller"
	"github.com/gorilla/mux"
	"net/http"
	"GP/constant"
	"GP/utils"
	"fmt"
)

func SetRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/",AllowOrigin(TokenCheck(keep))).Methods("GET")
	router.HandleFunc("/api/register", AllowOrigin(controller.Register)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/login", AllowOrigin(controller.Login)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/logout", AllowOrigin(TokenCheck(controller.LogOut))).Methods("GET", "OPTIONS")

	router.HandleFunc("/ws/chat", controller.WsMain)
	router.HandleFunc("/api/getonline", AllowOrigin(controller.GetOnline)).Methods("GET")

	router.HandleFunc("/api/getoneuser", AllowOrigin(controller.GetOneUser)).Methods("GET")
	router.HandleFunc("/api/getuserlist", AllowOrigin(controller.GetUserList)).Methods("GET")
	router.HandleFunc("/api/updateuser", AllowOrigin(controller.UpdateUser)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/updatepassword", AllowOrigin(controller.UpdatePassword)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/banuser", AllowOrigin(controller.BanUser)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/cancelbanuser", AllowOrigin(controller.CancelBanUser)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/upuserrole", AllowOrigin(controller.UpUserRole)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/downuserrole", AllowOrigin(controller.DownUserRole)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/finduser", AllowOrigin(controller.FindUser)).Methods("GET")
	router.HandleFunc("/api/getuserrole", AllowOrigin(controller.GetUserRole)).Methods("GET")

	router.HandleFunc("/api/getoneroom", AllowOrigin(controller.GetOneRoom)).Methods("GET")
	router.HandleFunc("/api/getroomlist", AllowOrigin(controller.GetRoomList)).Methods("GET")
	router.HandleFunc("/api/getuseroomlist", AllowOrigin(controller.GetUseRoomList)).Methods("GET")
	router.HandleFunc("/api/createroom", AllowOrigin(controller.CreateRoom)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/banroom", AllowOrigin(controller.BanRoom)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/cancelbanroom", AllowOrigin(controller.CancelBanRoom)).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/getcheckfriend", AllowOrigin(controller.GetCheckFriend)).Methods("GET")
	router.HandleFunc("/api/getfriendlist", AllowOrigin(controller.GetFriendList)).Methods("GET")
	router.HandleFunc("/api/newfriend", AllowOrigin(controller.NewFriend)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/passfriend", AllowOrigin(controller.PassFriend)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/unpassfriend", AllowOrigin(controller.UnPassFriend)).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/getcommentlist", AllowOrigin(controller.GetCommentList)).Methods("GET")
	router.HandleFunc("/api/newcomment", AllowOrigin(controller.NewComment)).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/gethistorylist", AllowOrigin(controller.GetHistoryList)).Methods("GET")

	return router
}

func keep(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	fmt.Println(r.Header)
	token := r.Header.Get("AccessToken")
	fmt.Println(token)
	info, err := utils.GetTokenInfo(token)
	fmt.Println(info)
	if err != nil {
		fmt.Println(err)
		return
	}
	fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbData(info).Response()
}

func AllowOrigin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Add("Access-Control-Allow-Headers", "AccessToken") //header的类型
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			return
		}
		next(w, r)
	})
}


