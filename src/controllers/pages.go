package controllers

import (
	"devbook/src/config"
	"devbook/src/cookies"
	"devbook/src/models"
	"devbook/src/requests"
	"devbook/src/responses"
	"devbook/src/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

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

	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteTemplate(w, "home.html", struct {
		Posts  []models.Post
		UserId uint64
	}{
		Posts:  posts,
		UserId: userId,
	})
}

func UpdatePostPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.API_URL, postId)
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

	var post models.Post
	if err = json.NewDecoder(response.Body).Decode(&post); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.Error{Error: err.Error()})
		return
	}
	utils.ExecuteTemplate(w, "updatePost.html", post)
}

func FilteredUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrUsername := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/users?filter=%s", config.API_URL, nameOrUsername)

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

	var users []models.User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.Error{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "users.html", users)
}

func UserPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	loggedUserId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userId == loggedUserId {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	userData, err := models.FindUserRelatedData(userId, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "user.html", struct {
		User         models.User
		LoggedUserId uint64
	}{
		User:         userData,
		LoggedUserId: loggedUserId,
	})
}

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	loggedUserId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	userData, err := models.FindUserRelatedData(loggedUserId, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "profile.html", userData)
}

func EditUserPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	loggedUserId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	channel := make(chan models.User)
	go models.FindUserData(channel, loggedUserId, r)
	user := <-channel

	if user.ID == 0 {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: "Error in finding user"})
		return
	}

	utils.ExecuteTemplate(w, "edit-user.html", user)
}

func UpdatePasswordPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "updatePassword.html", nil)
}
