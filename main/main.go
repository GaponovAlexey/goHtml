package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                 string
	Age                  uint16
	Money                int16
	AvgGrades, happiness float64
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("Username is:%s hi is age %d",
		u.Name, u.Age)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

var bob = User{"bob", 0, -20, 4.2, 0.9}

func main() {
	handlerRequest()
}

func handlerRequest() {
	fmt.Println("hi")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":3000", nil)
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go is super contacts page!")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/homePage.html")
	tmpl.Execute(w, bob)
}
