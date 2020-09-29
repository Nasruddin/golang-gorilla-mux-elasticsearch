package app

import (
	"github.com/nasruddin/golang/elastic/controllers"
	"net/http"
)

func mapUrls()  {
	router.HandleFunc("/ping", controllers.PingControllerInterface.Ping).Methods(http.MethodGet)
	router.HandleFunc("/users", controllers.UserControllerInterface.Create).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", controllers.UserControllerInterface.Get).Methods(http.MethodGet)
}
