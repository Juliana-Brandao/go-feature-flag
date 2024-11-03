package main

import (
	"database/sql"
	"github.com/Waelson/go-feature-flag/internal/controller"
	"github.com/Waelson/go-feature-flag/internal/repository"
	"github.com/Waelson/go-feature-flag/internal/service"
	"github.com/Waelson/go-feature-flag/internal/util"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {

	dbAddress := "localhost"
	if os.Getenv("ENV") == "container" {
		dbAddress = "postgres"
	}

	connStr := "postgres://myuser:mypassword@" + dbAddress + ":5432/mydb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	metricsRecord := util.NewMetricsRecord()
	featureFlagRepo := repository.NewFeatureFlagRepository(db)
	featureFlagService := service.NewFeatureFlagService(featureFlagRepo, metricsRecord)

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

	r.Handle("/metrics", promhttp.Handler())
	r.Get("/process-order", orderController.ProcessOrderHandler)

	log.Println("servidor inicialiado")
	http.ListenAndServe(":8080", r)
}
