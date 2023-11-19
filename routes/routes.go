package routes

import (
	"github.com/joao-antonio-gomes/web-mvc/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
