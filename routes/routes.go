package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/haakaashs/antino-labs/config"
)

func Start() {
	buildHandler()

	router := mux.NewRouter()
	router.HandleFunc("/students", basicHandler.GetStudents).Methods("GET")
	router.HandleFunc("/students", basicHandler.InsertOrUpdateStudents).Methods("POST")

	server := &http.Server{
		Addr:    config.Config.ServerPort,
		Handler: router,
	}

	log.Println("server started\nlistening and serving on port", config.Config.ServerPort)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Unable to start server", err)
	}

}
