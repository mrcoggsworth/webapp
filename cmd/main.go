package main

import (
	"fmt"
	handlers2 "github.com/mrcoggsworth/webapp/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers2.HomeHandler)
	http.HandleFunc("/about", handlers2.AboutHandler)

	fmt.Printf("Starting application on port %s\n", portNumber)
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		err.Error()
	}
}
