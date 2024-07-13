package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/garaevmir/simple-GOserver/server/users"
)

func AddNewUser(name, pass string) {
	ureq := users.UserRequest{
		Name:     name,
		Password: pass,
	}
	postdata, _ := json.Marshal(ureq)
	resp, err := http.Post("http://localhost:8080/user", "application/json", bytes.NewBuffer(postdata))
	if err != nil {
		println("Something went wrong")
		return
	}
	defer resp.Body.Close()
	fmt.Println("POST request response:", resp.Status)
	postRespBody := new(bytes.Buffer)
	n, _ := postRespBody.ReadFrom(resp.Body)
	fmt.Println("POST response body: ", postRespBody.String(), "Num of bytes: ", n)
}

func AddInfo(name, info string) {
	postdata := []byte("info: \"" + info + "\"")
	resp, err := http.Post("http://localhost:8080/user/"+name, "application/json", bytes.NewBuffer(postdata))
	if err != nil {
		println("Something went wrong")
		return
	}
	defer resp.Body.Close()
	fmt.Println("POST request response:", resp.Status)
	pos := new(bytes.Buffer)
	n, _ := pos.ReadFrom(resp.Body)
	fmt.Println("POST response body: ", pos.String(), "Num of bytes: ", n)
}

func DeleteUser(name string) {
	re, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/user/"+name, bytes.NewBuffer([]byte("")))
	if err != nil {
		println("Something went wrong")
		return
	}
	re.Header.Set("Content-Type", "application/json")
	reresp, err := http.DefaultClient.Do(re)
	if err != nil {
		println("Something went wrong")
		return
	}
	defer reresp.Body.Close()
	println(reresp.Status)
	postReBody := new(bytes.Buffer)
	n, _ := postReBody.ReadFrom(reresp.Body)
	fmt.Println("Delete response body: ", postReBody.String(), "Num of bytes: ", n)
}
