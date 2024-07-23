package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"product-image/files"
	"product-image/handlers"
	"time"

	gorillaHandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
)

func main() {

	l := hclog.New(
		&hclog.LoggerOptions{
			Name:  "product-image",
			Level: hclog.LevelFromString("loglevel"),
		},
	)

	// create a logger for the server from the default logger
	sl := l.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})

	if err := godotenv.Load(); err != nil {
		l.Error("Unable to locate .env file", err)
		os.Exit(1)
	}

	listenAddr := os.Getenv("LISTEN_ADDR")
	basePath := os.Getenv("BASE_PATH")
	if listenAddr == "" || basePath == "" {
		l.Error("Enviromental variable not set correctly")
		os.Exit(1)
	}

	// create a storage
	store, err := files.NewLocal(basePath)
	if err != nil {
		l.Error("Unable to create storage", err)
		os.Exit(1)
	}

	// create the handler
	fh := handlers.NewFiles(store, l)
	mw := handlers.GzipHandler{}

	// create a new serve mux
	sm := mux.NewRouter()

	ch := gorillaHandler.CORS(gorillaHandler.AllowedOrigins([]string{"http://localhost:5173"}))

	// filename regex: {filename:[a-zA-Z]+\\.[a-z]{3}}
	// problem with FileServer is that it is dumb
	ph := sm.Methods(http.MethodPost).Subrouter()
	ph.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.UploadREST)
	ph.HandleFunc("/", fh.UploadMultipart)

	// get files
	gh := sm.Methods(http.MethodGet).Subrouter()
	gh.Handle(
		"/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}",
		http.StripPrefix("/images/", http.FileServer(http.Dir(basePath))),
	)
	gh.Use(mw.GzipMiddleware)

	// create a new server
	s := http.Server{
		Addr:         listenAddr,        // configure the bind address
		Handler:      ch(sm),            // set the default handler
		ErrorLog:     sl,                // the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Info("Starting server", "bind_address", listenAddr)

		err := s.ListenAndServe()
		if err != nil {
			l.Error("Unable to start server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	l.Info("Shutting down server with", "signal", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
	defer cancel()
}
