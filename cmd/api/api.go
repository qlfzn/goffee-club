package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/qlfzn/goffee-club/internal/orders"
	"github.com/qlfzn/goffee-club/internal/products"
	"github.com/qlfzn/goffee-club/internal/repository"
)

type Application struct {
	Config Config
	DB     *pgxpool.Pool
	Logger *log.Logger
}

type Config struct {
	Addr string
	DB   *repository.DBConfig
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
		Addr:    app.Config.Addr,
		Handler: h,
	}

	app.Logger.Printf("server has started at http://localhost%s", app.Config.Addr)
	return srv.ListenAndServe()
}
