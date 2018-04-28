package main

import (
	"futureHealth/api"
	"futureHealth/lib"
	"strings"

	"encoding/json"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	achievPers := lib.CreateAchievmentPersistence("./data/achievments.json")
	achievHandler := lib.AchievmentHandler{&achievPers}

	userPersistence := lib.CreateUserPersistence("./data/users.json")
	runtasticApi := api.RuntasticApi{}
	userHandler := lib.UserHandler{&userPersistence, &runtasticApi}

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
		if token == "" {
			return
		}
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
		userIdBlob, found := c.Get("userId")
		if !found {
			c.JSON(401, "Authentication required")
		}
		achievments := achievHandler.Read()
		c.JSON(200, userHandler.UserAchieved(achievments, userIdBlob.(string)))
	})
	app.POST("/runtastic", func(c *gin.Context) {
		credentials := lib.Credentials{}
		bodyDecoder := json.NewDecoder(c.Request.Body)
		err := bodyDecoder.Decode(&credentials)
		if err != nil {
			c.JSON(400, "Invalid JSON")
			return
		}

		userIdBlob, found := c.Get("userId")
		if !found {
			c.JSON(401, "Authentication required")
		}

		err = userHandler.RuntasticLogin(credentials, userIdBlob.(string))
		if err != nil {
			c.JSON(400, "Invalid Login")
			return
		}
		c.JSON(200, "OK")
	})

	admin := r.Group("/admin")
	admin.GET("/achievments", func(c *gin.Context) {
		achievs := achievHandler.Read()
		c.JSON(200, achievs)
	})
	admin.POST("/achievment", func(c *gin.Context) {
		achiev := lib.Achievment{}
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
