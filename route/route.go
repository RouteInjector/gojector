package route

import "github.com/gin-gonic/gin"

// Route provides the definition of a URL route
type Route struct {
	Path   string // URL path where the route is map
	Method string // HTTP method the route uses
	Handler gin.HandlerFunc // Handler for the route
}