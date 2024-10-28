package ogen

import (
	"log"
	"net/http"
	h "server/presentation/http/ogen/handler"
	middlewares "server/presentation/http/ogen/middleware"
	oas "server/pkg/oas"

	"github.com/rs/cors"
)

func Init() {
	handler := h.NewHandler()

	server, err := oas.NewServer(handler, oas.WithMiddleware(middlewares.Logging()))
	if err != nil {
		log.Fatal(err)
	}
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// TODO: 環境変数に入れる
	if err := http.ListenAndServe(":18080", cors.Handler(server)); err != nil {
		log.Fatal(err)
	}
}
