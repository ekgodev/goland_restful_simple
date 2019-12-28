package main

import (
	"database/sql"
	"fmt"
	"github.com/ekprog/restful_test/src/database"
	_ "github.com/ekprog/restful_test/src/migrations"
	"log"
	"net/http"
	"strconv"
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

	// get id
	ids, ok := r.URL.Query()["id"]
	if !ok {
		http.Error(w, "GET parameter 'id' is required", http.StatusBadRequest)
		return
	}

	// parse if
	userIdStr := ids[0]
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		http.Error(w, "GET parameter 'id' is incorrect", http.StatusBadRequest)
		return
	}

	// database logic
	query := `SELECT user_name, phone FROM phones WHERE id=$1;`
	row := database.DB.QueryRow(query, userId)

	var userName, userPhone string
	err = row.Scan(&userName, &userPhone)
	switch err {
	case sql.ErrNoRows:
		http.Error(w, "Not found", http.StatusNotFound)
		return
	case nil:
	default:
		log.Println(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	ans := fmt.Sprintf("%s,%s", userName, userPhone)
	_, err = fmt.Fprintf(w, ans)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
}

// GET get_all
func getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// database logic
	query := `SELECT user_name, phone FROM phones;`
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	type userInfo struct {
		userName, userPhone string
	}

	var userInfos []userInfo
	for rows.Next() {
		var userName, userPhone string
		err = rows.Scan(&userName, &userPhone)
		if err != nil {
			log.Println(err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		userInfos = append(userInfos, userInfo{
			userName:  userName,
			userPhone: userPhone,
		})
	}

	ans := ""
	for i := range userInfos {
		ans += fmt.Sprintf("%s,%s\n", userInfos[i].userName, userInfos[i].userPhone)
	}

	_, err = fmt.Fprintf(w, ans)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
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
		Reload: false,
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
