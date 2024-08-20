package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/we-we-Web/draw-lots-backend/service"
)

func SetUpRouter(s *service.Service) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "NTOUCSE113")
	})
	router.GET("/api", func(c *gin.Context) {
		type Option struct {
			Admin  string
			Senior string
			Junior string
		}
		c.JSON(http.StatusOK, &Option{})
	})

	router.POST("/api/admin", s.CreateAdmin)
	router.GET("/api/admin/:id", s.GetAdmin)

	router.POST("/api/senior", s.CreateSenior)
	router.GET("api/seniors", s.GetAllSeniors)
	router.GET("/api/senior/:id", s.GetSenior)
	router.PUT("/api/senior/partner/:id", s.AddChildIdToSenior)

	router.POST("/api/junior", s.CreateJunior)
	router.GET("/api/juniors", s.GetAllJuniors)
	router.GET("/api/junior/:id", s.GetJunior)
	router.PUT("/api/junior/partner/:id", s.AddParentIdToJunior)
	router.PUT("/api/junior/line/:id", s.AddLineIdToJunior)

	return router
}
