package engine

import (
	"github.com/RouteInjector/gojector/server"
	"github.com/RouteInjector/gojector/infrastructure/mongo"
	"fmt"
)

type Engine struct {
	server *server.Server
	models []mongo.ModelWrapper
}

func NewEngine(server *server.Server) (*Engine) {
	return &Engine{
		server: server,
	}
}

func (e *Engine) CreateModelsREST(models []mongo.ModelWrapper) {
	e.models = models
	fmt.Println("CreateModelsREST")
	fmt.Println(e.models)
	for _, m := range e.models {
		fmt.Println("CreateModelsREST: " + m.Model.Name)
		if m.Model.Get {
			e.CreateGet(m)
		}
	}
}