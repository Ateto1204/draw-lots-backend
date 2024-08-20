package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/we-we-Web/draw-lots-backend/model"
)

func (service *Service) CreateJunior(c *gin.Context) {
	var junior model.Junior
	if err := c.ShouldBindJSON(&junior); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user, _ := service.juniorRepo.GetJunior(junior.StudentNumber); user != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}
	if err := service.juniorRepo.CreateJunior(&junior); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, junior)
}

func (service *Service) GetAllJunior(c *gin.Context) {

}

func (service *Service) GetJunior(c *gin.Context) {
	id := c.Param("id")
	junior, err := service.juniorRepo.GetJunior(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, junior)
}

func (service *Service) AddSeniorIdToJunior(c *gin.Context) {

}
