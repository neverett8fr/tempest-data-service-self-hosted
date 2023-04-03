package main

import (
	"log"
	"tempest-data-service/cmd"
	application "tempest-data-service/pkg/application/service"
	"tempest-data-service/pkg/config"

	"github.com/gorilla/mux"
)

// Route declaration
func getRoutes(conf config.Config) *mux.Router {
	r := mux.NewRouter()
	application.NewRoutes(r, conf)

	return r
}

// Initiate web server
func main() {
	conf, err := config.Initialise()
	if err != nil {
		log.Fatalf("error initialising config, err %v", err)
		return
	}
	log.Println("config initialised")

	// serviceDB, err := cmd.OpenDB(&conf.DB)
	// if err != nil {
	// 	log.Fatalf("error starting db, err %v", err)
	// 	return
	// }
	// defer serviceDB.Close()
	// log.Println("connection to DB setup")

	// err = cmd.MigrateDB(serviceDB, conf.DB.Driver)
	// if err != nil {
	// 	log.Fatalf("error running DB migrations, %v", err)
	// 	return
	// }
	// log.Println("DB migrations ran")

	router := getRoutes(*conf)
	log.Println("API routes retrieved")

	err = cmd.StartServer(&conf.Service, router)
	if err != nil {
		log.Fatalf("error starting server, %v", err)
		return
	}
	log.Println("server started")

}
