package server

import (
	"golang-rest-api-starter/controllers"
	"net/http"
)

func Init() {
	router := controllers.Endpoint
	http.ListenAndServe(":8080", router())
}
