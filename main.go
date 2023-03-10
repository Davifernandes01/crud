package main

import (
	metado "crud/metadosHttp"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/users", metado.Post).Methods(http.MethodPost)
	router.HandleFunc("/users", metado.Get).Methods(http.MethodGet)
	router.HandleFunc("/users/{Id}", metado.GetById).Methods(http.MethodGet)
	router.HandleFunc("/users/{Id}", metado.Update).Methods(http.MethodPut)
	router.HandleFunc("/users/{Id}", metado.Delete).Methods(http.MethodDelete)

	fmt.Println("ok")
	log.Fatal(http.ListenAndServe(":5000", router))
}
