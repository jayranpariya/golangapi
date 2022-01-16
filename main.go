package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jayranpariya/mongoapi/router"
)

func main() {
	fmt.Println("mongodb api")
	r := router.Router()
	fmt.Println("server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ")

}
