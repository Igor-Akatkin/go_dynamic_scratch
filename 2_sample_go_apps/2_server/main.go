package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/users", usersPage)

	port := ":9091"
	println("Server listen on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsFired   bool
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	// user := User{"CoolName", "CoolLastName"}
	// js, _ := json.Marshal(user)

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func usersPage(w http.ResponseWriter, r *http.Request) {
	users := []User{User{"Max", "Power", false}, User{"John", "Wheels", true}}
	// js, _ := json.Marshal(user)

	tmpl, err := template.ParseFiles("static/users.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
