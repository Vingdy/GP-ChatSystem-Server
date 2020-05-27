// main.go文件
package main

import (
	"GP/controller"
	"GP/db"
	"GP/redis"
	"GP/router"
	"net/http"
	"fmt"
)

func main() {
	db.InitMySql()
	db.InitTable()
	redis.InitRedis()
	controller.Ws_init()
	err := http.ListenAndServe(":9000", router.SetRouter())
	fmt.Println("Listening...")
	if err == nil {
		fmt.Println(err)
	}

}
