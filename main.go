package main

import (
	"futureHealth/achievment"
	"futureHealth/user"
	"strings"

	"encoding/json"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	achievPers := achievment.CreateJsonPersistence("./data/achievments.json")
	achievHandler := achievment.AchievmentHandler{&achievPers}

	userPersistence := user.CreateJsonPersistence("./data/users.json")
	userHandler := user.UserHandler{&userPersistence}

	r := gin.Default()

	r.Use(cors.Default())

	app := r.Group("/app")
	app.POST("/login", func(c *gin.Context) {
		c.JSON(200, userHandler.Create())
	})
	app.GET("/achievments", func(c *gin.Context) {
		achievs := achievHandler.Read()
		c.JSON(200, achievs)
	})

	app.Use(func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		token := strings.TrimPrefix(bearerToken, "Bearer ")
		c.Set("userId", token)
	})
	app.GET("/points", func(c *gin.Context) {
		userIdBlob, found := c.Get("userId")
		if !found {
			c.JSON(401, "Authentication required")
		}
		c.JSON(200, userHandler.Points(userIdBlob.(string)))
	})
	app.GET("/achieved", func(c *gin.Context) {
		c.JSON(200, "/achieved")
	})

	admin := r.Group("/admin")
	admin.GET("/achievments", func(c *gin.Context) {
		achievs := achievHandler.Read()
		c.JSON(200, achievs)
	})
	admin.POST("/achievment", func(c *gin.Context) {
		achiev := achievment.Achievment{}
		bodyDecoder := json.NewDecoder(c.Request.Body)
		err := bodyDecoder.Decode(&achiev)
		if err != nil {
			c.JSON(400, "Invalid JSON")
			return
		}

		achievHandler.Create(achiev)
		c.JSON(200, "OK")
	})
	r.Run(":8000")
}
