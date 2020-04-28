// main.go文件
package main

import (
	"GP/db"
	"GP/router"
	"log"
	"net/http"
	"GP/redis"
)

func main() {

	db.InitMySql()
	db.InitTable()
	redis.InitRedis()

	err := http.ListenAndServe(":8080", router.SetRouter())
	if err == nil {
		log.Println(err)
	}

}
