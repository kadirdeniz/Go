package main

import (
	"fmt"
	"net/http"
	_ "github.com/gorm/database"
	"github.com/gorm/database/model"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func createUser(w http.ResponseWriter,r *http.Request){
	var user user.User = user.User{
		Name:"Kadir",
		Surname:"Deniz",
		Email:"kadirdenz@hotmail.com",
	}
	user.Create()
	fmt.Fprintf(w,"User Created")
}

func deleteUser(w http.ResponseWriter,r *http.Request){
	var user user.User = user.User{
		ID:1,
	}
	user.Delete()
	fmt.Fprintf(w,"User Deleted")
}

func getUser(w http.ResponseWriter,r *http.Request){
	var user user.User = user.User{
		ID:1,
	}
	user.Get()
	fmt.Fprintf(w,"User Created")
}

func handleRequests() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/create",createUser)
	http.HandleFunc("/delete",deleteUser)
	http.HandleFunc("/get",getUser)

	http.ListenAndServe(":8080",nil)
}


func main() {
	fmt.Println("Server Started")

	handleRequests()
}
