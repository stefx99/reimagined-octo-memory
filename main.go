package main

import (
	"IS_project/internal/controller"
	"fmt"

	"log"
	"net/http"
)

func main() {
	r := controller.Router()

	fmt.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
