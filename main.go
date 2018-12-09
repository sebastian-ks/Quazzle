package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type config struct {
	Port string `json:"port"`
}

type userType struct {
	Name string
}

var user userType

func main() {
	conf := getConfig()

	//Static files like css and js
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	http.Handle("/script/", http.StripPrefix("/script/", http.FileServer(http.Dir("script"))))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/register", handleRegister)

	http.ListenAndServe(conf.Port, nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	if user.Name == "" {
		user.Name = "blabla"
	}
	templ.Execute(w, user)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("login.html")
		if err != nil {
			panic(err)
		}
		templ.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		user.Name = r.Form["username"][0]
	}

}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("register.html")
	if err != nil {
		panic(err)
	}
	templ.Execute(w, nil)
}

func getConfig() config {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	var conf config
	json.Unmarshal([]byte(data), &conf)
	return conf
}
