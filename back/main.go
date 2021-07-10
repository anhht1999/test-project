package main

import (
	"log"
	"net/http"
	"back/database"
	"back/routers"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	r := mux.NewRouter()
	routers.Setup(r)

	log.Fatal(http.ListenAndServe(":3000", r))
}