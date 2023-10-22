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

	// user APIs
	router.HandleFunc("/user", basicHandler.GetStudents).Methods("POST")
	router.HandleFunc("/students", basicHandler.InsertOrUpdateStudents).Methods("POST")

	//products APIs
	
	// cart APIs

	// order APIs
















	// server creation with port specified
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
