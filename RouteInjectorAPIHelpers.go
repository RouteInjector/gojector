package routeinjector

import (
	"github.com/RouteInjector/gojector/model"
	"github.com/RouteInjector/gojector/route"
)

func (r *RouteInjector) AddModel(model *model.Model) {
	if r.started {
		glog.Fatal("Can not add a new model when server is started")
	}
	r.Models = append(r.Models, *r.mongo.WrapModel(model))
}

func (r *RouteInjector) AddRoute(route route.Route) {
	if r.started {
		glog.Fatal("Can not add a new route when server is started")
	}
	r.Routes = append(r.Routes, route)
}