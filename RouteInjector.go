package routeinjector

import (
	"github.com/RouteInjector/gojector/infrastructure/conf"
	"github.com/op/go-logging"
	"github.com/RouteInjector/gojector/server"
	"github.com/RouteInjector/gojector/model"
	"github.com/RouteInjector/gojector/route"
)

var glog = logging.MustGetLogger("RouteInjector")

type RouteInjector struct {
	Config  *conf.Configuration
	Models  []model.Model
	Routes  []route.Route
	server  *server.Server
	started bool
}

func NewRouteInjector(config *conf.Configuration) *RouteInjector {
	glog.Info("      __________               __           ")
	glog.Info("      \\______   \\ ____  __ ___/  |_  ____   ")
	glog.Info("       |       _//  _ \\|  |  \\   __\\/ __ \\  ")
	glog.Info("       |    |   (  <_> )  |  /|  | \\  ___/  ")
	glog.Info("       |____|_  /\\____/|____/ |__|  \\___  > ")
	glog.Info("              \\/                        \\/  ")
	glog.Info(" .___            __               __                  ")
	glog.Info(" |   | ____     |__| ____   _____/  |_  ___________   ")
	glog.Info(" |   |/    \\    |  |/ __ \\_/ ___\\   __\\/  _ \\_  __ \\  ")
	glog.Info(" |   |   |  \\   |  \\  ___/\\  \\___|  | (  <_> )  | \\/  ")
	glog.Info(" |___|___|  /\\__|  |\\___  >\\___  >__|  \\____/|__|     ")
	glog.Info("          \\/\\______|    \\/     \\/                     ")
	glog.Info("\n")
	r := &RouteInjector{}
	r.Config = config
	r.server = server.NewServer(r.Config)
	return r
}

func (r *RouteInjector) Start() {
	glog.Debug("Starting RouteInjector instance")
	r.server.InjectModels(r.Models)
	r.server.InjectRoutes(r.Routes)
	r.server.Start()
	r.started = true
}