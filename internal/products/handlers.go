package products

import (
	"encoding/json"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler{
	return &handler{
		service: service,
	}
}

func (h *handler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}