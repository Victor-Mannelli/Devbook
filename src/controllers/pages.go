package controllers

import (
	"devbook/src/config"
	"devbook/src/models"
	"devbook/src/requests"
	"devbook/src/responses"
	"devbook/src/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func SingUpPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.API_URL)

	response, err := requests.ReqWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorHandler(w, response)
		return
	}

	var posts []models.Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.Error{Error: err.Error()})
		return
	}
	// fmt.Println(posts)
	utils.ExecuteTemplate(w, "home.html", posts)
}
