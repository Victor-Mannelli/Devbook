package controllers

import (
	"devbook/src/cookies"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Clear(w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
