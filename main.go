package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/jaspersurmont/rsloggers-api/api"
	"github.com/jaspersurmont/rsloggers-api/storage"
	"github.com/joho/godotenv"
)

func main() {
  slog.SetLogLoggerLevel(slog.LevelDebug.Level())

  // Load env variables
  err := godotenv.Load()
  if err != nil {
    slog.Error("could not load .env file %w", err)
    os.Exit(1)
  }

  storeProvider, err := storage.NewStoreProvider()
  if err != nil {
    slog.Error(err.Error())
    os.Exit(1)
  }

	router := http.NewServeMux()
	humaApi := humago.New(router, huma.DefaultConfig("RS Loggers", "0.1"))
	api.Setup(humaApi, storeProvider)


  slog.Info("listening on port 8081")
	http.ListenAndServe("127.0.0.1:8081", router)
}
