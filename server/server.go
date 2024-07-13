package main

import (
	"fmt"
	"net/http"

	"github.com/garaevmir/simple-GOserver/server/users"
)

var amount int
var DB users.UsersStorage

func main() {
	DB = users.NewUsersStorage()
	http.HandleFunc("/user", HandleUser)
	http.HandleFunc("/user/{name}", HandleUserName)
	http.HandleFunc("/{$}", HandleStart)
	http.HandleFunc("/", HandleAll)
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
