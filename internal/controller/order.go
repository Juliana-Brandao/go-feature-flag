package controller

import (
	"github.com/Waelson/go-feature-flag/internal/service"
	"net/http"
)

import (
	"github.com/go-chi/chi/v5"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (oc *OrderController) ProcessOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "orderID")

	result, err := oc.orderService.ProcessOrder(orderID)
	if err != nil {
		http.Error(w, "Erro ao processar pedido", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(result))
}
