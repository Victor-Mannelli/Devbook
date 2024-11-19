package models

import (
	"devbook/src/config"
	"devbook/src/requests"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	Followers []User    `json:"followers"`
	Following []User    `json:"following"`
	Posts     []Post    `json:"posts"`
}

func FindUserRelatedData(userId uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	postChannel := make(chan []Post)

	go FindUserData(userChannel, userId, r)
	go FindFollowers(followersChannel, userId, r)
	go FindFollowing(followingChannel, userId, r)
	go FindPost(postChannel, userId, r)

	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case receivedUser := <-userChannel:
			if receivedUser.ID == 0 {
				return User{}, errors.New("error in finding user")
			}

			user = receivedUser

		case receivedFollowers := <-followersChannel:
			if receivedFollowers == nil {
				return User{}, errors.New("error in finding followers")
			}

			followers = receivedFollowers

		case receivedFollowing := <-followingChannel:
			if receivedFollowing == nil {
				return User{}, errors.New("error in finding followed users")
			}

			following = receivedFollowing

		case receivedPosts := <-postChannel:
			if receivedPosts == nil {
				return User{}, errors.New("error in finding posts")
			}

			posts = receivedPosts
		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

func FindUserData(channel chan<- User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.API_URL, userId)

	response, err := requests.ReqWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}
func FindFollowers(channel chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/followers/%d", config.API_URL, userId)

	response, err := requests.ReqWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err = json.NewDecoder(response.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followers
}
func FindFollowing(channel chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/followers/%d/following", config.API_URL, userId)

	response, err := requests.ReqWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if err = json.NewDecoder(response.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	if following == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- following
}
func FindPost(channel chan<- []Post, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/posts/user/%d", config.API_URL, userId)

	response, err := requests.ReqWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		channel <- nil
		return
	}

	if posts == nil {
		channel <- make([]Post, 0)
		return
	}

	channel <- posts
}
