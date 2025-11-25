package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/qlfzn/goffee-club/internal/orders"
)

type Application struct {
	Addr string
	// db
	// logger
}

// create HTTP router for api
func (app *Application) Mount() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})
	r.Post("/orders", func(w http.ResponseWriter, r *http.Request) {
		var order orders.NewOrder
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// call order service
		orderDetails := &orders.OrderDetails{}
		checkOrder, err := orderDetails.CreateNewOrder(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// return response to client
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Received order!"))
		err = json.NewEncoder(w).Encode(&checkOrder)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	return r
}

// run server
func (app *Application) Run(h http.Handler) error {
	srv := &http.Server{
		Addr:    app.Addr,
		Handler: h,
	}

	log.Printf("server has started at http://localhost%s", app.Addr)
	return srv.ListenAndServe()
}
