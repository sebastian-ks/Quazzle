package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Config struct {
	Port string `json:"port"`
}

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
	templ.Execute(w, nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("login.html")
	if err != nil {
		panic(err)
	}
	templ.Execute(w, nil)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("register.html")
	if err != nil {
		panic(err)
	}
	templ.Execute(w, nil)
}

func getConfig() Config {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()
	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var conf Config
	json.Unmarshal([]byte(data), &conf)
	return conf
}
