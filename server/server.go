package server

import (
	"github.com/op/go-logging"
	"github.com/gin-gonic/gin"
	"github.com/RouteInjector/gojector/infrastructure/conf"
	"strconv"
	"net/http"
	"strings"
	"github.com/RouteInjector/gojector/route"
	"github.com/RouteInjector/gojector/infrastructure/mongo"
)

var glog = logging.MustGetLogger("server")

// Server struct
type Server struct {
	Port   int
	router *gin.Engine
}

func NewServer(conf *conf.Configuration) *Server {
	server := &Server{
		Port: conf.Bind,
		router: gin.Default(),
	}
	server.version()
	return server
}

func (s *Server) version() {
	s.router.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "TakaWacala")
	})
}

func (s *Server) InjectModels(wrappers []mongo.ModelWrapper) {
	for _, wrap := range wrappers {
		s.router.GET("/" + strings.ToLower(wrap.Model.Name), test(wrap.Model.Name))
	}
}

func (s *Server) InjectRoutes(routes []route.Route) {
	for _, route := range routes {
		s.router.Handle(route.Method, route.Path, route.Handler)
	}
}

func (s *Server) Start() {
	s.router.Run(":" + strconv.Itoa(s.Port))
}

func test(modelName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, modelName)
	}
}