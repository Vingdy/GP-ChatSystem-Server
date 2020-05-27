package router

import (
	"GP/constant"
	"GP/controller"
	"GP/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var path, err = utils.GetProDir()

var FileDownloadHost="/static/"

var FileDownloadPrefix="/files"

var TemplateDir =  path+"/dist/"

func SetRouter() *mux.Router {
	router := mux.NewRouter()

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../dist/angular"))))
	filesever:=http.StripPrefix(FileDownloadHost, http.FileServer(http.Dir(path+FileDownloadPrefix)))

	router.PathPrefix(FileDownloadHost).HandlerFunc(filesever.ServeHTTP)

	router.PathPrefix("/static").Handler(http.StripPrefix(
		"/static",http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		http.FileServer(http.Dir(TemplateDir+`angular`)).ServeHTTP(w,r)
	})))
	router.PathPrefix("/assets").Handler(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		http.FileServer(http.Dir(TemplateDir+`angular`)).ServeHTTP(w,r)
	}))
	router.NotFoundHandler =  http.HandlerFunc(notFound)


	router.HandleFunc("/test", AllowOrigin(TokenCheck(keep))).Methods("GET")
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

func notFound(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,path+`/dist/angular`)
}

func keep(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	fmt.Println(r.Header)
	token := r.Header.Get("AccessToken")
	fmt.Println(token)
	fmt.Println(path)
	fb.FbCode(constant.SUCCESS).FbMsg(token).FbData(path).Response()
}

func AllowOrigin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Add("Access-Control-Allow-Headers", "AccessToken")  //header的类型
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			return
		}
		next(w, r)
	})
}
