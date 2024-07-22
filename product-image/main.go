package main

import (
	"os"
	"product-image/files"
	"product-image/handlers"

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
}
