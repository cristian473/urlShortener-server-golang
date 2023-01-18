package main

import (
	"fmt"
)

func main(){
	server := NewServer(":8080")
	server.Handle("/", MiddlewareRoot)
	server.Handle("/home", MiddlewareHome)
	err := server.Listen()
	if err != nil {
		fmt.Println(err)
	}
}