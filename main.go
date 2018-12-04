package main

import (
	"html/template"
	"net/http"
)

func main() {
	//Static files like css and js
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	http.Handle("/script/", http.StripPrefix("/script/", http.FileServer(http.Dir("script"))))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/register", handleRegister)

	http.ListenAndServe(":3001", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	templ, _ := template.ParseFiles("index.html")
	templ.Execute(w, nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	templ, _ := template.ParseFiles("login.html")
	templ.Execute(w, nil)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	templ, _ := template.ParseFiles("register.html")
	templ.Execute(w, nil)
}
