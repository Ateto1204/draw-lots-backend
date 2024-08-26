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

	// router.POST("/api/admin", s.CreateAdmin)
	router.POST("/api/admin/:id", s.GetAdminById)

	// router.POST("/api/senior", s.CreateSenior)
	router.POST("api/seniors", s.GetAllSeniors)
	router.POST("/api/senior/:id", s.GetSeniorById)

	// router.POST("/api/junior", s.CreateJunior)
	router.POST("/api/juniors", s.GetAllJuniors)
	router.POST("/api/junior/:id", s.GetJuniorById)
	router.PUT("/api/junior/line/:id", s.AddLineIdToJunior)

	return router
}
