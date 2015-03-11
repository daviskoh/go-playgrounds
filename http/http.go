package main

import (
	"fmt"
	"io"
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	// log the requset for illustrative purposes
	fmt.Println("********************* req *********************")
	fmt.Println(req)
	fmt.Println("********************* req *********************")

	// write to res
	// res.Write([]byte("meow"))

	io.WriteString(res, "meow")
}

func main() {
	http.HandleFunc("/bro", Handler)
	http.HandleFunc("/hello", Handler)
	http.ListenAndServe(":3000", nil)
}
