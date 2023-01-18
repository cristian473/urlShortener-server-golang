package main

import (
	"fmt"
	"net/http"
)

func MiddlewareRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func MiddlewareHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "home middleware")
}