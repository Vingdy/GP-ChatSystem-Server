// main.go文件
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Here is the home page."))
	})

	http.ListenAndServe(":8080", nil)
}