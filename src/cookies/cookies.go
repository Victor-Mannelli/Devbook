package cookies

import (
	"devbook/src/config"
	"net/http"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Config() {
	s = securecookie.New(config.HASH_KEY, config.BLOCK_KEY)
}

func Save(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	hashedData, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    hashedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}
