package main

import (
	// "SvGorm/models"

	"SvGorm/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	// models.MigrateUser2()
	mux.HandleFunc("/api/v2/user/", handlers.GetUsersV2).Methods("GET")
	mux.HandleFunc("/api/v2/user/{id:[0-9]+}", handlers.GetUserV2).Methods("GET")
	mux.HandleFunc("/api/v2/user/", handlers.SaveUser).Methods("POST")

	// mux.HandleFunc("/api/user", handlers.GetUsers).Methods("GET")
	// mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUserId).Methods("GET")
	// mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	// mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.EditUser).Methods("PUT")
	// mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3300", mux))
}
