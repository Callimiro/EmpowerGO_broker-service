package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}

	/* 	r := chi.NewRouter()
	   	r.Use(middleware.Logger)
	   	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	   		w.Write([]byte("welcome"))
	   	})
	   	http.ListenAndServe(":3000", r) */
}
