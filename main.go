// main.go文件
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", newPage)
	err := http.ListenAndServe(":80",nil)
	if err != nil {
		log.Fatal(err)
	}
}

func newPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is New Page!")
}
