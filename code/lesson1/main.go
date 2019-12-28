package main

import (
	"log"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello World!!!")
}

func main() {

	// 127.0.0.1:8080/test
	http.HandleFunc("/test", test)

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
	    panic(err)
	}
}