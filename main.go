// main.go文件
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Here is the home page."))
	})

	http.ListenAndServe(":80", nil)
}
