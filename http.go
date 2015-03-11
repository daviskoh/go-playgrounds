package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	// log the requset
	fmt.Println("req")
	fmt.Println(req)

	// write to res
	// res.Write([]byte("meow"))

	io.WriteString(res, "meow")
}

func main() {
	http.HandleFunc("/bro", handler)
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":3000", nil)
}
