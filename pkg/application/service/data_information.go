package service

import (
	"net/http"
	application "tempest-data-service/pkg/application/entities"

	"github.com/gorilla/mux"
)

func newDataInformation(r *mux.Router) {
	r.HandleFunc("/test/{text}", testHandler).Methods(http.MethodGet)
	r.HandleFunc("/data/{username}", userFileNames).Methods(http.MethodGet)

}

func userFileNames(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	usr := params[username]

	fileNames, err := StorageProvider.GetAllFileInformation(r.Context(), usr)
	if err != nil {
		body := application.NewResponse(nil, err)
		writeReponse(w, body)
		return
	}

	body := application.NewResponse(fileNames)
	writeReponse(w, body)
}
