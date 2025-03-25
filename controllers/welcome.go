package controllers

import (
	"github.com/gin-gonic/gin"
)

type Welcome_controller struct{}

func (n *Welcome_controller) Init_Welcome_controller(route *gin.Engine){
	welcome := route.Group("/welcome")
	welcome.GET("/", n.GetWelcome())
	welcome.POST("/", n.PostWelcome())
}

func (n *Welcome_controller) GetWelcome() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(200,gin.H{
			"Hii":"welcome get request",
		})
	}
}


func (n *Welcome_controller) PostWelcome() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(200,gin.H{
			"Hii":"welcome post request",
		})
	}
}