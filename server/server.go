package server

import (
	"github.com/op/go-logging"
	"github.com/gin-gonic/gin"
	"github.com/RouteInjector/gojector/infrastructure/conf"
	"strconv"
	"net/http"
	"github.com/RouteInjector/gojector/route"
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

func (s *Server) GET(path string, handler gin.HandlerFunc){
	s.router.GET(path, handler)
}

func (s *Server) POST(path string, handler gin.HandlerFunc){
	s.router.POST(path, handler)
}

func (s *Server) PUT(path string, handler gin.HandlerFunc){
	s.router.PUT(path, handler)
}

func (s *Server) DELETE(path string, handler gin.HandlerFunc){
	s.router.DELETE(path, handler)
}

func (s *Server) InjectRoutes(routes []route.Route) {
	for _, route := range routes {
		s.router.Handle(route.Method, route.Path, route.Handler)
	}
}

func (s *Server) Start() {
	s.router.Run(":" + strconv.Itoa(s.Port))
}