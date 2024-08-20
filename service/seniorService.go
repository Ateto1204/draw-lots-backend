package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/we-we-Web/draw-lots-backend/model"
)

func (service *Service) CreateSenior(c *gin.Context) {
	var senior model.Senior
	if err := c.ShouldBindJSON(&senior); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user, _ := service.seniorRepo.GetSenior(senior.StudentNumber); user != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}
	if err := service.seniorRepo.CreateSenior(&senior); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, senior)
}

func (service *Service) GetAllSeniors(c *gin.Context) {
	seniors, err := service.seniorRepo.GetAllSeniors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		return
	}
	c.JSON(http.StatusOK, seniors)
}

func (service *Service) GetSenior(c *gin.Context) {
	id := c.Param("id")
	senior, err := service.seniorRepo.GetSenior(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, senior)
}

func (service *Service) AddJuniorIdToSenior(c *gin.Context) {

}
