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
	AvgGrades, Happiness float64
	Hobbies              []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he mobey equel: %d", u.Name, u.Age, u.Money)
}
func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"Football", "Skate", "Dance"}}
	/*	bob.setNewName("Alex")
		fmt.Fprintf(w, bob.getAllInfo()) */
	temp, _ := template.ParseFiles("templates/homePage.html")
	temp.Execute(w, bob)
}
func contats_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contacts")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contats_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
