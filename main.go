// main.go文件
package main

import (
	"GP/db"
	"GP/router"
	"log"
	"net/http"
	"GP/redis"
	"GP/controller"
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
