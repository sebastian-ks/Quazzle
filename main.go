package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
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
	ID       int
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
	if user.Name == "" {
		user.Name = "blabla"
	}
	handleTempl("index.html", w, user)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handleTempl("login.html", w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		err := db.QueryRow("SELECT PASSWORD FROM user WHERE USERNAME = ?", r.Form["username"][0]).Scan(&user.Password) //DB naming conventions still to be set
		if err != nil {
			if err == sql.ErrNoRows {
				data := struct {
					ErrMsg string
				}{"Username not found or no input at all"}
				handleTempl("login.html", w, data)
			} else {
				panic(err)
			}
		} else
		//Have to put rows.Close after checkErr because Close would fail in case of an error
		if isPasswordCorrect(r.Form["password"][0]) {
			fmt.Println("Validation succesfull")
			user.Name = r.Form["username"][0]
			http.Redirect(w, r, "index", http.StatusSeeOther)
		} else {
			data := struct {
				ErrMsg string
			}{"Invalid input"}
			handleTempl("login.html", w, data)
		}
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handleTempl("register.html", w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form["displayname"])
		statement, err := db.Prepare("INSERT INTO user(USERNAME, PASSWORD) VALUES(?, ?)")
		checkErr(err)
		statement.Exec(r.Form["displayname"][0], hashPW(r.Form["password"][0]))
		user.Name = r.Form["displayname"][0]
		http.Redirect(w,r,"/",http.StatusSeeOther)
	}
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

func handleTempl(page string, w io.Writer, data interface{}) {
	templ, err := template.ParseFiles(page)
	if err != nil {
		panic(err)
	}

	templ.Execute(w, data)
}

func isPasswordCorrect(formPW string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formPW)); err != nil {
		fmt.Println("Error in password verification: ", err)
		return false
	}
	return true
}
