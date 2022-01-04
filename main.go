package main

import (
	"blog-mongo/app/config"
	"blog-mongo/app/setup"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {

	setup.AuthRouter()
	setup.ArticleRouter()
	setup.CommentRouter()

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		Debug:          true,
	}).Handler(setup.Mux)
	err := http.ListenAndServe(":"+config.DefaultConfig["appPort"], handler)
	if err != nil {
		log.Fatal(err)
	}
}
