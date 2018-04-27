package main

import (
	"futureHealth/business"

	"encoding/json"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	persistence := business.CreateJsonPersistence("./data/achievments.json")
	handler := business.AchievmentHandler{&persistence}

	r := gin.Default()

	r.Use(cors.Default())

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
		c.JSON(200, handler.Read())
	})
	admin.POST("/achievment", func(c *gin.Context) {
		achiev := business.Achievment{}
		bodyDecoder := json.NewDecoder(c.Request.Body)
		err := bodyDecoder.Decode(&achiev)
		if err != nil {
			c.JSON(400, "Invalid JSON")
			return
		}

		handler.Create(achiev)
		c.JSON(200, "OK")
	})
	r.Run(":8000")
}
