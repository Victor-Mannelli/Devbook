package controllers

import (
	"bytes"
	"devbook/src/config"
	"devbook/src/cookies"
	"devbook/src/requests"
	"devbook/src/responses"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"username": r.FormValue("username"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.API_URL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorHandler(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func Unfollow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/followers/%d", config.API_URL, userId)
	response, err := requests.ReqWithAuth(r, http.MethodDelete, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorHandler(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func Follow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/followers/%d", config.API_URL, userId)
	response, err := requests.ReqWithAuth(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorHandler(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"username": r.FormValue("username"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.API_URL, userId)
	response, err := requests.ReqWithAuth(r, http.MethodPut, url, bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorHandler(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	passwords, err := json.Marshal(map[string]string{
		"password":    r.FormValue("password"),
		"newPassword": r.FormValue("newPassword"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/changePassword", config.API_URL)
	response, err := requests.ReqWithAuth(r, http.MethodPost, url, bytes.NewBuffer(passwords))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorHandler(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
