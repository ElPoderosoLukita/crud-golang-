package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ElPoderosoLukita/gocrud/models"
	"github.com/gorilla/mux"
)

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	user := models.NewUser()
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, models.GenerateResponseJson(http.StatusBadRequest, err))
	}

	err = models.PostUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, models.GenerateResponseJson(http.StatusInternalServerError, err))
	}

	fmt.Fprint(w, models.GenerateResponseJson(http.StatusCreated, "User created correctly."))
}

func PutUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	user := models.NewUser()
	idInt, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, models.GenerateResponseJson(http.StatusBadRequest, err))
	}

	err = models.PutUser(user, idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, models.GenerateResponseJson(http.StatusInternalServerError, err))
	}

	fmt.Fprint(w, models.GenerateResponseJson(http.StatusOK, "User updated correctly."))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idInt, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := models.DeleteUser(idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, models.GenerateResponseJson(http.StatusInternalServerError, err))
	}

	fmt.Fprint(w, models.GenerateResponseJson(http.StatusOK, "User deleted correctly."))

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)

	users := models.GetUsers()

	fmt.Fprint(w, models.GenerateResponseJson(http.StatusOK, users))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)

	idInt, _ := strconv.Atoi(mux.Vars(r)["id"])
	user := models.GetUser(idInt)

	fmt.Fprint(w, models.GenerateResponseJson(http.StatusOK, user))
}
