package controllers

import (
	"dev-book/src/database"
	"dev-book/src/models"
	"dev-book/src/repositories"
	"dev-book/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.CreateUsersRepository(db)
	userRepository.Create(user)

	responses.JSON(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.CreateUsersRepository(db)
	users, err := userRepository.GetUsers()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	vars := mux.Vars(r)

	userRepository := repositories.CreateUsersRepository(db)
	user, err := userRepository.GetUser(vars["userId"])

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	vars := mux.Vars(r)
	userId := vars["userId"]

	userRepository := repositories.CreateUsersRepository(db)
	user, err := userRepository.GetUser(userId)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var userData models.User
	if err = json.Unmarshal(body, &userData); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	user, err = userRepository.Update(userId, user, userData)

	responses.JSON(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	vars := mux.Vars(r)
	userId := vars["userId"]

	userRepository := repositories.CreateUsersRepository(db)

	err = userRepository.Delete(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
