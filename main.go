package main

import (
	"log"
	"net/http"

	"github.com/HIldebrandoFerreira/goapi/router"
)

func main() {
	rt := router.Routers()
	address := "127.0.0.1:8081"
	log.Printf("Running api restfull on : http://%s\n", address)
	log.Fatal(http.ListenAndServe(address, rt))

}
