package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/fusco2k/go-web/exercices/hands-on07/01/models"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

type UserController struct {
	storeDB map[int]models.User
}

func NewUserController(db map[int]models.User) *UserController {
	return &UserController{db}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		log.Fatalln(err)
	}

	if uc.storeDB[id].Id != id {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// composite literal
	u := models.User{}

	// Fetch user
	u = uc.storeDB[id]

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	_, err = fmt.Fprintf(w, "%s\n", uj)
	if err != nil {
		log.Fatal(err)
	}
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	if uc.storeDB[u.Id].Id == u.Id {
		_, err := fmt.Fprint(w, "user already exists")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	uc.storeDB[u.Id] = u

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		log.Fatalln(err)
	}

	if uc.storeDB[id].Id != id {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	delete(uc.storeDB, id)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")
}
