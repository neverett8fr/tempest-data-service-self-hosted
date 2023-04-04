package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"tempest-data-service/pkg/config"
	"tempest-data-service/pkg/infra/storage"

	"github.com/gorilla/mux"
)

const (
	username = "username"
	item     = "item"

	headerAccept = "Accept"

	headerContentType = "Content-Type"
	contentTypeJSON   = "application/json"
	// contentTypeImage  = "image/"
)

var (
	StorageProvider storage.StorageProvider
)

func NewRoutes(r *mux.Router, conf config.Config) {
	sp, err := storage.InitialiseStorageProvider(
		context.Background(),
		conf.Storage.MountLocation,
	)
	if err != nil {
		log.Printf("error initialising storage provider, err %v", err)
	}

	StorageProvider = sp

	newGeneric(r)
	newDataInformation(r)
	newDataOperation(r)
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

func writeFile(w http.ResponseWriter, body []byte) {

	// Set the response headers
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", fmt.Sprintf("%v", len(body)))

	_, err := w.Write(body)
	if err != nil {
		log.Printf("error writing response, err %v", err)
	}
}
