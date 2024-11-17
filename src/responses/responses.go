package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	} else {
		w.Write([]byte("{}"))
	}
}

func ErrorHandler(w http.ResponseWriter, r *http.Response) {
	var err Error
	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)
}
