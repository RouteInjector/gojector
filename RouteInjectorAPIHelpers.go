package routeinjector

import (
	"github.com/RouteInjector/gojector/model"
	"github.com/RouteInjector/gojector/route"
)

func (r *RouteInjector) AddModels(models ...model.Model) {
	if r.started {
		glog.Fatal("Can not add a new model when server is started")
	}
	r.Models = append(r.Models, models...)
}

func (r *RouteInjector) AddModel(model model.Model) {
	if r.started {
		glog.Fatal("Can not add a new model when server is started")
	}
	r.Models = append(r.Models, model)
}

func (r *RouteInjector) AddRoutes(route ...route.Route) {
	if r.started {
		glog.Fatal("Can not add a new route when server is started")
	}
	r.Routes = append(r.Routes, route...)
}

func (r *RouteInjector) AddRoute(route route.Route) {
	if r.started {
		glog.Fatal("Can not add a new route when server is started")
	}
	r.Routes = append(r.Routes, route)
}