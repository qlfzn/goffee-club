package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/qlfzn/goffee-club/internal/orders"
	"github.com/qlfzn/goffee-club/internal/products"
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
	
	// declare services & handlers
	orderService := orders.NewService()
	orderHandler := orders.NewHandler(orderService)
	r.Post("/orders", orderHandler.CreateNewOrder)

	productService := products.NewService()
	productHandler := products.NewHandler(productService)
	r.Get("/products", productHandler.GetProductsHandler)

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
