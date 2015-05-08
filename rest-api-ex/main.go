package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/daviskoh/go-playgrounds/rest-api-ex/burgers"
	"github.com/julienschmidt/httprouter"
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

	server := negroni.Classic()
	server.UseHandler(router)

	server.Run(":3000")
}
