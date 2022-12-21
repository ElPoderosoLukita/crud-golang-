package main

import (
	"net/http"

	"github.com/ElPoderosoLukita/gocrud/handlers"
	"github.com/ElPoderosoLukita/gocrud/models"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	models.OpenDB()
	defer models.CloseDB()

	r.HandleFunc("/post/user", handlers.PostUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/put/user/{id:[0-9]+}", handlers.PutUserHandler).Methods(http.MethodPut)
	r.HandleFunc("/delete/user/{id:[0-9]+}", handlers.DeleteUserHandler).Methods(http.MethodDelete)
	r.HandleFunc("/get/users", handlers.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/get/user/{id:[0-9]+}", handlers.GetUser).Methods(http.MethodGet)

	http.ListenAndServe(":8081", r)
}
