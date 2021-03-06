package main

import (
	"encoding/json"
	"futureHealth/api"
	"futureHealth/lib"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	achievPers := lib.CreateAchievmentPersistence("./data/achievments.json")
	achievHandler := lib.AchievmentHandler{&achievPers}

	stressPers := lib.CreateStressPersistence("./data/stress.json")
	stressHandler := lib.StressLevelHandler{&stressPers}

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
	app.GET("/achievements", func(c *gin.Context) {
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
	app.POST("/stress", func(c *gin.Context) {
		stress := lib.StressLevel{}
		bodyDecoder := json.NewDecoder(c.Request.Body)
		err := bodyDecoder.Decode(&stress)
		if err != nil {
			c.JSON(400, "Invalid JSON")
			return
		}
		userIdBlob, found := c.Get("userId")
		if !found {
			c.JSON(401, "Authentication required")
		}
		stressHandler.Create(stress, userIdBlob.(string))
		c.JSON(200, "OK")
	})
	app.GET("/stress", func(c *gin.Context) {
		stress := stressHandler.Read()
		c.JSON(200, stress)
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

	r.Run(":3000")
}
