package main

import (
	"fmt"
	"github.com/ekprog/restful_test/src/database"
	_ "github.com/ekprog/restful_test/src/migrations"
	"log"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request URI: " + r.RequestURI)
	fmt.Println("Request method: " + r.Method)

	fmt.Println("Params: ")
	for k, v := range r.URL.Query() {
		fmt.Println(k + " = " + v[0])
	}
}

// GET /get?id=2
func get(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// GET get_all
func getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// UPDATE /update?id=2&phone=79185555555
func update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "UPDATE" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// PUT /add?id=2&phone=79185555555
func add(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func main() {

	// DATABASE
	dbSettings := database.Settings{
		User:   "postgres",
		Pass:   "",
		Host:   "localhost",
		Port:   "5432",
		Name:   "phone_list",
		Reload: true,
	}
	err := database.Connect(dbSettings)
	if err != nil {
		log.Fatal(err)
	}

	// ROUTER
	http.HandleFunc("/test", test)
	http.HandleFunc("/get", get)
	http.HandleFunc("/get_all", getAll)
	http.HandleFunc("/update", update)
	http.HandleFunc("/add", add)

	fmt.Printf("Запуск сервера\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
