package main

import (
	"log"
	"log/slog"
	"net/http"
	"ordersystem/repository"
	"ordersystem/rest"
	"ordersystem/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	_ "ordersystem/docs"

	// OpenApi
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title				Order System
// @description			This system enables drink orders and should not be used for the forbidden Hungover Games.
func main() {
	// connect to s3
	s3, err := storage.CreateS3client()
	if err != nil {
		log.Fatalln(err)
	}
	// connect to db
	db, err := repository.NewDatabaseHandler()
	if err != nil {
		log.Fatalln(err)
	}
	// prepopulate data
	err = repository.Prepopulate(db, s3)
	if err != nil {
		log.Fatalln(err)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// allow local cors
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Origin", "cache-control", "expires", "pragma"},
		ExposedHeaders:   []string{"Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Menu Routes
	r.Get("/api/menu", rest.GetMenu(db))
	// Order Routes
	r.Get("/api/order/all", rest.GetOrders(db))
	r.Get("/api/order/totalled", rest.GetOrdersTotal(db))
	r.Get("/api/receipt/{orderId}", rest.GetReceiptFile(db, s3))
	r.Post("/api/order", rest.PostOrder(db, s3))
	// OpenAPI Routes
	r.Get("/openapi/*", httpSwagger.WrapHandler)

	slog.Info("⚡⚡⚡ Order System is up and running ⚡⚡⚡")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}
