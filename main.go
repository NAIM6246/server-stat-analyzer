package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/naim6246/server-stat-analyzer/serverStat"
)

func main() {
	router := chi.NewRouter()
	serverStat := serverStat.NewServerStat()

	router.Get("/server-stat",func(w http.ResponseWriter, r *http.Request) {
		serverStat.ServerStat()
	})
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080",router)
}