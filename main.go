package main

import (
	"github.com/gin-gonic/gin"
	"github.com/niteshKrr/gin-framework/config"
	"github.com/niteshKrr/gin-framework/routes"
)

func main() {
	
	config.Connect_DB()
	config.MigrateDB()
	router := gin.Default()

	user_routes := routes.User_routes{}
	user_routes.Init_user_routes(router)

	router.Run(":8000")
}
