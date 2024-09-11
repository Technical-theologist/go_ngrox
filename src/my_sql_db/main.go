package main

import (
	"fmt"
	"log"
	_ "mydatabase/controllers"
	"mydatabase/routers"
	"net/http"
)

func main() {

	fmt.Println("MangoDB API")
	r := routers.Router()
	fmt.Println("Server is getting  started")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 .....")
}
