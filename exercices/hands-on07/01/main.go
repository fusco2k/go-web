package main

import (
	"github.com/fusco2k/go-web/exercices/hands-on07/01/controllers"
	"github.com/fusco2k/go-web/exercices/hands-on07/01/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[int]models.User {
	// Connect to our local mongo
	storedb := map[int]models.User{}

	return storedb
}
