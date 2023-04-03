package service

import (
	"encoding/json"
	"log"
	"net/http"

	"tempest-data-service/pkg/infra/storage"

	"github.com/gorilla/mux"
)

const (
	username = "username"
	item     = "item"
)

var (
	StorageProvider storage.StorageProvider
)

func NewRoutes(r *mux.Router) {
	newGeneric(r)

	r.HandleFunc("/test/{text}", testHandler).Methods("GET")
}

func writeReponse(w http.ResponseWriter, body interface{}) {

	reponseBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("error converting reponse to bytes, err %v", err)
	}
	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write(reponseBody)
	if err != nil {
		log.Printf("error writing response, err %v", err)
	}
}
