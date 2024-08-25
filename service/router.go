package service

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(s *Service) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

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

	router.POST("/api/login", s.Login)
	router.PUT("/api/connect", s.CreateConnect)

	router.POST("/api/admin", s.CreateAdmin)
	router.GET("/api/admin/:id", s.GetAdminById)

	router.POST("/api/senior", s.CreateSenior)
	router.GET("api/seniors", s.GetAllSeniors)
	router.GET("/api/senior/:id", s.GetSeniorById)

	router.POST("/api/junior", s.CreateJunior)
	router.GET("/api/juniors", s.GetAllJuniors)
	router.GET("/api/junior/:id", s.GetJuniorById)
	router.PUT("/api/junior/line/:id", s.AddLineIdToJunior)

	return router
}
