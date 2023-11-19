package main

import (
	"fmt"
	"github.com/joao-antonio-gomes/web-mvc/routes"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	fmt.Println("Servidor rodando no endereço http://localhost:8000")
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
