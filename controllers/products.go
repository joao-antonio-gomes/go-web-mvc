package controllers

import (
	"github.com/joao-antonio-gomes/web-mvc/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.FindAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}
