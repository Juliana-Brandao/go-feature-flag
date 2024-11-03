package controller

import (
	"github.com/Waelson/go-feature-flag/internal/service"
	"net/http"
	"strconv"
)

type FeatureFlagController struct {
	service *service.FeatureFlagService
}

func NewFeatureFlagController(service *service.FeatureFlagService) *FeatureFlagController {
	return &FeatureFlagController{service: service}
}

func (controller *FeatureFlagController) UpdateFeatureFlagStatusHandler(w http.ResponseWriter, r *http.Request) {
	flagName := r.URL.Query().Get("flag_name")
	valueStr := r.URL.Query().Get("value")

	if flagName == "" || valueStr == "" {
		http.Error(w, "Os parâmetros 'flag_name' e 'value' são obrigatórios", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		http.Error(w, "Valor inválido para o parâmetro 'value'", http.StatusBadRequest)
		return
	}

	if err := controller.service.UpdateFeatureFlagStatus(flagName, value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Feature flag atualizada com sucesso"))
}
