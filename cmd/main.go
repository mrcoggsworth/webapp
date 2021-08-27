package main

import (
	"fmt"
	"github.com/mrcoggsworth/webapp/pkg/config"
	"github.com/mrcoggsworth/webapp/pkg/handlers"
	"github.com/mrcoggsworth/webapp/pkg/utils"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := utils.CreateTemplateCache()
	if err != nil {
		log.Fatalln("[!] cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	utils.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.HomeHandler)
	http.HandleFunc("/about", handlers.Repo.AboutHandler)

	fmt.Printf("Starting application on port %s\n", portNumber)
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		err.Error()
	}
}
