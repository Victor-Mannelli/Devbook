package controllers

import (
	"devbook/src/config"
	"devbook/src/requests"
	"devbook/src/utils"
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

	fmt.Println(response.StatusCode, err)

	utils.ExecuteTemplate(w, "home.html", nil)
}
