package engine

import (
	"github.com/RouteInjector/gojector/infrastructure/mongo"
	"github.com/gin-gonic/gin"
	"fmt"
)

func (e *Engine) CreateGet(m mongo.ModelWrapper) {
	fmt.Println("Create GET: " +m.Model.GetPathName() + "/:id")
	e.server.GET(m.Model.GetPathName() + "/:id", getHandler(m))
}

func getHandler(m mongo.ModelWrapper) gin.HandlerFunc {
	return func(c *gin.Context) {
		doc, _ := m.FindOne(c.Param("id"))
		c.JSON(200, doc)
	}
}