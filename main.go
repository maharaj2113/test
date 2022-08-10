package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maharaj2113/test/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listing to Port 4000")
}
