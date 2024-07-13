package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/garaevmir/simple-GOserver/server/users"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	amount++
	println(amount)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte(DB.GetAll()))
	case http.MethodPost:
		buf := new(bytes.Buffer)
		if n, err := buf.ReadFrom(r.Body); err != nil || n == 0 {
			w.Write([]byte("Empty request or undefined error"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var ureq users.UserRequest
		if err := json.Unmarshal(buf.Bytes(), &ureq); err != nil {
			w.Write([]byte("Deserialization error"))
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		if DB.FindUser(ureq.Name) {
			w.Write([]byte("Already exists"))
			w.WriteHeader(http.StatusConflict)
			return
		}
		user := users.NewUser(ureq.Name, ureq.Password)
		DB.AddUser(&user)
		println("Here: ", buf.String())
		w.Write([]byte(user.String() + "\n"))
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func HandleUserName(w http.ResponseWriter, r *http.Request) {
	amount++
	println(amount)
	switch r.Method {
	case http.MethodGet:
		var name string
		fmt.Sscanf(r.RequestURI, "/user/%s", &name)
		if DB.FindUser(name) {
			user, _ := DB.GetUser(name)
			w.Write([]byte(user.String() + "\n"))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	case http.MethodPost:
		var name string
		fmt.Sscanf(r.RequestURI, "/user/%s", &name)
		if !DB.FindUser(name) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		user, _ := DB.GetUser(name)
		buf := new(bytes.Buffer)
		if _, err := buf.ReadFrom(r.Body); err != nil {
			w.Write([]byte("Empty request or undefined error"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		println(buf.String())
		var temp string
		fmt.Sscanf(buf.String(), "info: %q", &temp)
		user.Info = temp
		println(user.Info)
		w.Write([]byte("Info updated"))
	case http.MethodDelete:
		var name string
		fmt.Sscanf(r.RequestURI, "/user/%s", &name)
		DB.DeleteUser(name)
		w.Write([]byte("User: " + name + " deleted\n"))
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func HandleStart(w http.ResponseWriter, r *http.Request) {
	amount++
	println(amount)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello world!"))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func HandleAll(w http.ResponseWriter, r *http.Request) {
	amount++
	println(amount)
	w.WriteHeader(http.StatusNotFound)
}
