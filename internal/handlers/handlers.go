package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func Home(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	templateSet, err := template.ParseFiles("./internal/templates/home.tmpl.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(res, "Internal Server Error", 500)
		return
	}

	err = templateSet.Execute(res, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(res, "Internal Server Error", 500)
		return
	}
}

