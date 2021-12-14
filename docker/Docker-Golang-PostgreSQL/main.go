package main

import (
	"fmt"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PUser, PPassword, PDb, PPort, PHost, Port string = os.Getenv("POSTGRES_USER"),
	os.Getenv("POSTGRES_PASSWORD"),
	os.Getenv("POSTGRES_DB"),
	os.Getenv("POSTGRES_PORT"),
	os.Getenv("POSTGRES_HOST"),
	os.Getenv("PORT")

var DSN string = "host=" + PHost + " user=" + PUser + " password=" + PPassword + " dbname=" + PDb + " port=" + PPort + " sslmode=disable TimeZone=Asia/Istanbul"

func main() {
	_, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Db Success")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Anasayfa")
	})

	http.ListenAndServe(":8091", nil)

}
