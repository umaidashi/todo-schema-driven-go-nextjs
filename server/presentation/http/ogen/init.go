package ogen_server

import (
	"log"
	"net/http"
	h "server/presentation/http/ogen/handler"
	oas "server/presentation/http/ogen/oas"

	"github.com/rs/cors"
)

func Init() {
	handler := h.NewHandler()

	server, err := oas.NewServer(handler)
	if err != nil {
		log.Fatal(err)
	}
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// 環境変数を入れる
	if err := http.ListenAndServe(":18080", cors.Handler(server)); err != nil {
		log.Fatal(err)
	}
}
