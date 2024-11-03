package service

import (
	"fmt"
)

type OrderService struct {
	featureFlagService *FeatureFlagService
}

func NewOrderService(featureFlagService *FeatureFlagService) *OrderService {
	return &OrderService{featureFlagService: featureFlagService}
}

func (os *OrderService) ProcessOrder(orderID string) (string, error) {
	// Verifica se a feature flag "processOrderFeature" está ativa
	if os.featureFlagService.IsFeatureEnabled("processOrderFeature") {
		// Nova funcionalidade de processamento do pedido
		result := fmt.Sprintf("Nova funcionalidade: Pedido %s processado com sucesso!", orderID)
		return result, nil
	}

	// Funcionalidade padrão de processamento do pedido
	result := fmt.Sprintf("Funcionalidade padrão: Pedido %s processado com sucesso!", orderID)
	return result, nil
}
