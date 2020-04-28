package router

import (
	"GP/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func SetRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/",AllowOrigin(TokenCheck(keep))).Methods("GET")
	router.HandleFunc("/api/register", AllowOrigin(controller.Register)).Methods("POST", "OPTION")
	router.HandleFunc("/api/login", AllowOrigin(controller.Login)).Methods("POST", "OPTION")
	router.HandleFunc("/ws/chat", controller.WsMain)

	return router
}

func keep(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here is the home page."))
}

func AllowOrigin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Add("Access-Control-Allow-Origin",  "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	})
}


