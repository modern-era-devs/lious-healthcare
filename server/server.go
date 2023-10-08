package server

import (
	"github.com/gorilla/mux"
	"github.com/modern-era-devs/lious-healthcare/handlers"
	"log"
	"net/http"
)

type apiServer struct {
	Host string
	Port string
}

func StartAPIServer() error {
	router := getRouter()

	// TODO move these to configs
	s := apiServer{
		Host: "",
		Port: "3000",
	}

	log.Println("host: ", s.Host, "port: ", s.Port)
	if err := http.ListenAndServe(s.Host+":"+s.Port, router); err != nil && err != http.ErrServerClosed {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func getRouter() *mux.Router {
	router := mux.NewRouter()

	// system routes
	router.HandleFunc("/ping", handlers.PingHandler).Methods(http.MethodOptions, http.MethodGet)

	// user routes

	return router
}
