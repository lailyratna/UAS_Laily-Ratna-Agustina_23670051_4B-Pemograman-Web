package main

import (
	"fmt"
	"net/http"

	"sistem-peminjaman-lab/config"
	"sistem-peminjaman-lab/routes"
)

func main() {
	config.ConnectDB()

	r := routes.SetupRoutes()

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
