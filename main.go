package main

import (
	"log"
	"net/http"

	"github.com/ulugbekrozimboyev/rest-api-server/utils"
)

func main() {

	router := utils.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
