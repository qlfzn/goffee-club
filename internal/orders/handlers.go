package orders

import (
	"encoding/json"
	"net/http"
)

// create http handler that:
//	- returns a handler with service
// 	- receives input & validate
//	- send response to client

type NewOrder struct {
	ItemName string  `json:"itemName"`
	Size     string  `json:"size"`
	Price    string `json:"price"`
}

type OrderDetails struct {
	Name string
	Size string
	Price float64
	Status string
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler{
	return &handler{
		service: service,
	}
}

func (h *handler) CreateNewOrder(w http.ResponseWriter, r *http.Request) {
		var order NewOrder
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// call order service
		checkOrder, err := h.service.CreateNewOrder(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// return response to client
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Received order!"))
		err = json.NewEncoder(w).Encode(&checkOrder)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
}