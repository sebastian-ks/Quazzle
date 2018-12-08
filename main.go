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

	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("index.html")
	if err != nil { panic(err) }
	templ.Execute(w, nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("login.html")
	if err != nil { panic(err) }
	templ.Execute(w, nil)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("register.html")
	if err != nil { panic(err) }
	templ.Execute(w, nil)
}
