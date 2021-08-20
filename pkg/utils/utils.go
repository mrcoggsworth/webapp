package utils

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		log.Println(fmt.Errorf("template parse error: %v", err))
		return
	}

	if err = parsedTemplate.Execute(w, nil); err != nil {
		log.Println(fmt.Errorf("error parsing template: %v", err))
		return
	}
}
