package main

import (
	"github.com/RouteInjector/gojector/infrastructure/conf"
	"github.com/RouteInjector/gojector"
	"github.com/RouteInjector/gojector/model"
	"github.com/RouteInjector/gojector/route"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	model.Schema
	Name string `json:"name",bson:"name"`
}

func main() {
	conf := &conf.Configuration{
		Database: &conf.Database{
			Endpoint: "localhost:27017",
			Name:"gojector",
		},
	}
	conf.Auth = false
	conf.Bind = 40000

	personModel := model.Model{
		Name: "Person",
		Plural: "Persons",
		ID: "name",
		Schema: Person{},
		Get: true,
		Post:true,
	}

	route := route.Route{
		Method: "GET",
		Path: "/taka",
		Handler: func(c *gin.Context) {
			c.String(http.StatusOK, "Ok, it is working")
		},
	}

	injector := routeinjector.NewRouteInjector(conf)
	injector.AddModel(&personModel)
	injector.AddRoute(route)
	injector.Start()
}