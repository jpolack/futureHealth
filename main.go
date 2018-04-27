package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	app := r.Group("/app")
	app.POST("/login", func(c *gin.Context) {
		c.JSON(200, "/login")
	})
	app.GET("/achievments", func(c *gin.Context) {
		c.JSON(200, "/achievments")
	})
	app.GET("/points", func(c *gin.Context) {
		c.JSON(200, "/points")
	})
	app.GET("/achieved", func(c *gin.Context) {
		c.JSON(200, "/achieved")
	})

	admin := r.Group("/admin")
	admin.GET("/achievments", func(c *gin.Context) {
		c.JSON(200, "/achievments")
	})
	admin.POST("/achievments", func(c *gin.Context) {
		c.JSON(200, "/achievments")
	})
	r.Run(":8000")
}
