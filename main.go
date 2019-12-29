// main.go文件
package main

import (
	"GP/db"
	"GP/router"
	"log"
	"net/http"
)

func main() {

	db.InitMySql()
	db.InitTable()

	err := http.ListenAndServe(":80", router.SetRouter())
	if err == nil {
		log.Println(err)
	}
}
