package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pfirulo2022/go-htmx-crud/database"
)

func main() {
	database.ConnectDB()

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs)) //

	fmt.Println("Running in port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
