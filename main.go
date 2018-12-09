package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type config struct {
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type userType struct {
	Name     string
	Id       int
	Seshid   int
	Password string
}

var user userType

var db *sql.DB

func main() {
	conf := getConfig()

	db = initDBConn(conf)
	defer db.Close()

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

		rows, err := db.Query("SELECT hashed_pw FROM users WHERE name=\"" + r.Form["username"][0] + "\"")
		checkErr(err)
		rows.Next()
		rows.Scan(&user.Password)
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Form["password"][0])); err != nil {
			fmt.Println(err, "Validation unsuccesfull")
		} else {
			fmt.Println("Validation succesfull")
		}

	}

}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("register.html")
	if err != nil {
		panic(err)
	}
	templ.Execute(w, nil)
}

func initDBConn(conf config) *sql.DB {
	database, err := sql.Open("mysql", conf.Username+":"+conf.Password+"@tcp(localhost:3306)/quazzle")
	checkErr(err)

	err = database.Ping()
	checkErr(err)
	return database
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

func hashPW(pw string) []byte {
	fmt.Println([]byte(pw))
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	checkErr(err)
	return hash
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
