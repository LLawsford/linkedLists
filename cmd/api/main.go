package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

type config struct {
	port int
	env  string
}

const version = "1.0.0"

type application struct {
	config config
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 3030, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production)")
	flag.Parse()

	app := &application{config: cfg}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	srv.ListenAndServe()
}
