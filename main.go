package main

import (
	"github.com/joao-antonio-gomes/web-mvc/routes"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	log.Println("Servidor rodando no endereço http://localhost:8000")
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
