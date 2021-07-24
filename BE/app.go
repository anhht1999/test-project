package main

import (
	"fmt"
	"log"
	"net/http"
	"ocg-be/database"
	"ocg-be/routes"
	
	"ocg-be/rmq-sendmail/reportService"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func main() {
	db, err := database.Connect() //Khởi tạo biến conection
	if err != nil {               //Catch error trong quá trình thực thi
		log.Printf("Error %s when getting db connection", err)
		return
	}
	defer db.Close()
	log.Printf("Successfully connected to database")

	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public"))
    router.PathPrefix("/public").Handler(http.StripPrefix("/public", fs))

	handleCross := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"http://localhost:8080"}),
		handlers.AllowCredentials(),
	)

	go func() {
		reportService.RunReportService()
	}()
	routes.Setup(router)
	fmt.Println("Server running at localhost:8000")

	http.ListenAndServe(":8000", handleCross(router))

}
