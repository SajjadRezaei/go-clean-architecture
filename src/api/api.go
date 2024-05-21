package api

import (
	"github.com/SajjadRezaei/go-clean-architecture/api/routes"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1/")
	{
		//v1.GET("/health", func(c *gin.Context) {
		//	c.JSON(http.StatusOK, "Working!")
		//	return
		//})

		health := v1.Group("/health")
		routes.Health(health)

	}

	r.Run(":5005")
}
