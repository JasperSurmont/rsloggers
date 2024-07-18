package main

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/jaspersurmont/rsloggers-api/api"
)

func main() {
	router := http.NewServeMux()
	humaApi := humago.New(router, huma.DefaultConfig("RS Loggers", "0.1"))
	api.Setup(humaApi)

	http.ListenAndServe("127.0.0.1:8081", router)
}
