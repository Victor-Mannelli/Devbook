package controllers

import (
	"devbook/src/utils"
	"net/http"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}
