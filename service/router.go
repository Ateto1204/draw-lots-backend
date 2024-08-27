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

	router.POST("/api/login", s.Login)          // identity, id, pwd
	router.PUT("/api/connect", s.CreateConnect) // id, pwd, parent, child
	router.PUT("/api/clear", s.ClearConnection) // id, pwd

	router.POST("/api/admin", s.GetAdminById) // id, pwd

	router.GET("/api/senior/:id", s.GetSeniorById)
	router.GET("/api/junior/:id", s.GetJuniorById)

	router.PUT("/api/junior/line", s.SetLineIdToJunior) // id, pwd, line

	// router.POST("/api/admin", s.CreateAdmin)
	// router.POST("/api/senior", s.CreateSenior)
	// router.POST("api/seniors", s.GetAllSeniors)
	// router.POST("/api/junior", s.CreateJunior)
	// router.POST("/api/juniors", s.GetAllJuniors)

	return router
}
