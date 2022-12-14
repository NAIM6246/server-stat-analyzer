package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/naim6246/server-stat-analyzer/configs"
	"github.com/naim6246/server-stat-analyzer/serverStat"
)

func main() {
	config := configs.GetAppConfig()
	router := chi.NewRouter()
	serverStat := serverStat.NewServerStat()
	go serverStat.StoreStats()

	//cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/current-server-stat", func(w http.ResponseWriter, r *http.Request) {
		stat := serverStat.GetCurrentServerStat()

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(stat)
	})

	router.Get("/server-stats", func(w http.ResponseWriter, r *http.Request) {
		stats := serverStat.GetLoggedServerStat()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(stats)
	})
	fmt.Println("Server running on port: ", config.ListenPort)
	http.ListenAndServe(fmt.Sprintf(":%d", config.ListenPort), router)
}
