package main

import (
	"devbook/src/router"
	"devbook/src/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	utils.LoadTemplates()
	r := router.NewRouter()

	fmt.Println("Listening port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
