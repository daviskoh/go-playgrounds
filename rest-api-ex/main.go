package main

import (
	"fmt"
	"github.com/daviskoh/go-playgrounds/rest-api-ex/burgers"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()

	router.GET("/", Index)

	router.GET("/burgers", burgers.Index)

	router.GET("/burgers/:id", burgers.Show)

	fmt.Println("listening on port :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
