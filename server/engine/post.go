package engine

import (
	"github.com/RouteInjector/gojector/infrastructure/mongo"
	"github.com/gin-gonic/gin"
	"fmt"
	"reflect"
	"net/http"
)

func (e *Engine) CreatePost(m mongo.ModelWrapper) {
	fmt.Println("Create POST: " + m.Model.GetPathName())
	e.server.POST(m.Model.GetPathName(), getPostHandler(m))
}

func getPostHandler(m mongo.ModelWrapper) gin.HandlerFunc {
	mType := reflect.TypeOf(m.Model.Schema)
	fmt.Println(mType)  // prints "main.t1"
	return func(c *gin.Context) {

		newInstance := reflect.New(mType).Elem().Interface()
		if err := c.BindJSON(&newInstance); err == nil {
			fmt.Println("Printing instance")
			fmt.Println(reflect.TypeOf(newInstance).(reflect.Type))
			fmt.Println(newInstance)
			if err = m.Insert(newInstance); err == nil {
				c.JSON(http.StatusCreated, newInstance)
			} else {
				panic(err)
				c.JSON(http.StatusInternalServerError, err)

			}
		} else {
			panic(err)
			c.JSON(http.StatusInternalServerError, err)
		}
	}
}