package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/naim6246/server-stat-analyzer/serverStat"
)

func main() {
	router := chi.NewRouter()
	serverStat := serverStat.NewServerStat()

	router.Get("/server-stat",func(w http.ResponseWriter, r *http.Request) {
		stat := serverStat.ServerStat()
		
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(stat)
	})
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080",router)
}