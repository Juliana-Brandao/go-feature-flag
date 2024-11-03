package main

import (
	"database/sql"
	"github.com/Waelson/go-feature-flag/internal/controller"
	"github.com/Waelson/go-feature-flag/internal/repository"
	"github.com/Waelson/go-feature-flag/internal/service"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	featureFlagRepo := repository.NewFeatureFlagRepository(db)
	featureFlagService := service.NewFeatureFlagService(featureFlagRepo)

	if err := featureFlagService.UpdateFeatureFlags(); err != nil {
		log.Fatal("Erro ao atualizar feature flags:", err)
	}

	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for range ticker.C {
			if err := featureFlagService.UpdateFeatureFlags(); err != nil {
				log.Println("Erro ao atualizar feature flags:", err)
			} else {
				log.Println("Feature flags atualizadas.")
			}
		}
	}()

	orderService := service.NewOrderService(featureFlagService)
	orderController := controller.NewOrderController(orderService)

	r := chi.NewRouter()

	r.Get("/process-order", orderController.ProcessOrderHandler)

	http.ListenAndServe(":8080", r)
}
