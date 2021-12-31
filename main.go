package main

import (
	"blog-mongo/app/config"
	"blog-mongo/app/setup"
	"log"
	"net/http"
)

func main() {

	setup.AuthRouter()
	setup.ArticleRouter()
	setup.CommentRouter()

	err := http.ListenAndServe(":"+config.DefaultConfig["appPort"],setup.Mux)
	if err != nil {
		log.Fatal(err)
	}
}

