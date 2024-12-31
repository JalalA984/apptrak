package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func ping(res http.ResponseWriter, _ *http.Request) {
	res.Write([]byte("OK"))
}

func Home(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	files := []string{
		"./internal/templates/base.tmpl.html",
		"./internal/templates/components/navbar.tmpl.html",
		"./internal/templates/home.tmpl.html",
	}

	templateSet, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(res, "Internal Server Error", 500)
		return
	}

	err = templateSet.ExecuteTemplate(res, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(res, "Internal Server Error", 500)
		return
	}
}

func Login(res http.ResponseWriter, req *http.Request) {
	files := []string{
		"./internal/templates/base.tmpl.html",
		"./internal/templates/login.tmpl.html",
	}

	templateSet, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(res, "Internal Server Error", 500)
		return
	}

	err = templateSet.ExecuteTemplate(res, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(res, "Internal Server Error", 500)
		return
	}
}

func Register(res http.ResponseWriter, req *http.Request) {
	files := []string{
		"./internal/templates/base.tmpl.html",
		"./internal/templates/register.tmpl.html",
	}

	templateSet, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(res, "Internal Server Error", 500)
		return
	}

	err = templateSet.ExecuteTemplate(res, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(res, "Internal Server Error", 500)
		return
	}
}
