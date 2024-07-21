package api

import (
	"log/slog"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jaspersurmont/rsloggers-api/storage"
)

func Setup(api huma.API, storeProvider *storage.StoreProvider) {
  api.UseMiddleware(logApiCalls)
  newPlayerController(api, storeProvider.PlayerStore) 
}

func logApiCalls(ctx huma.Context, next func(huma.Context)) {
  slog.Info("Handling API call", "method", ctx.Method(), "path", ctx.URL().Path)
  next(ctx)
}
