package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niteshKrr/gin-framework/controllers"
)

func main() {
	router := gin.Default()

	Welcome_controller := controllers.Welcome_controller{}
	Welcome_controller.Init_Welcome_controller(router)

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// The request responds to an url matching: /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.POST("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action"
		c.String(http.StatusOK, "%t", b)
	})

	router.DELETE("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id": "deleted id is "+ id,
		})
	})

	router.Run(":8000")
}
