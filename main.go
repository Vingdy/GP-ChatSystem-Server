// main.go文件
package main

import (
	"GP/controller"
	"GP/db"
	"GP/redis"
	"GP/router"
	"log"
	"net/http"
)

func main() {

	db.InitMySql()
	db.InitTable()
	redis.InitRedis()
	controller.Ws_init()
	err := http.ListenAndServe(":8080", router.SetRouter())
	if err == nil {
		log.Println(err)
	}

}
