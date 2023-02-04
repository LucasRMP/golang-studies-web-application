package main

import (
	"fmt"
	"net/http"

	"github.com/lucasrmp/web-application-studies/routes"
)

func main() {
	routes.RegisterRoutes()
	fmt.Println("Server is listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
