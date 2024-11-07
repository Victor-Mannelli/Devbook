package main

import (
	"devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server running at port 3000")

	r := router.Setup()
	log.Fatal(http.ListenAndServe(":3000", r))
}
