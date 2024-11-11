package main

import (
	"devbook/src/config"
	"devbook/src/cookies"
	"devbook/src/router"
	"devbook/src/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	cookies.Config()

	utils.LoadTemplates()
	r := router.NewRouter()

	fmt.Printf("Listening port %d\n", config.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), r))
}
