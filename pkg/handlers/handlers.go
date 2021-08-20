package handlers

import (
	"fmt"
	utils2 "github.com/mrcoggsworth/webapp/pkg/utils"
	"log"
	"net/http"
)

// HomeHandler is a function to run when route hits /
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils2.RenderTemplate(w, "home.page.tmpl")
	log.Printf(fmt.Sprintf("The %v page was visited...", r.URL.Path))
}

// AboutHandler is a function to run when route hits /about
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	utils2.RenderTemplate(w, "about.page.tmpl")
	log.Printf(fmt.Sprintf("The %v page was visited...", r.URL.Path))
}

