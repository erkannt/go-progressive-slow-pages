package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", landingHandler)
	router.GET("/baseline", baselineHandler)
	router.GET("/chunked", chunkedHandler)
	router.GET("/chunked-templ", chunkedWithTemplHandler)
	router.GET("/slots", slotsHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
