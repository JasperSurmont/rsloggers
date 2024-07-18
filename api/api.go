package api

import (
	"github.com/danielgtaylor/huma/v2"
)

type Controller struct {
	api huma.API
}

func Setup(api huma.API) Controller {
	c := Controller{api: api}
	c.setupPlayer()

	return c
}
