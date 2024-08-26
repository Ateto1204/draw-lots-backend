package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/we-we-Web/draw-lots-backend/model"
)

func (service *Service) CreateAdmin(c *gin.Context) {
	var admin model.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user, _ := service.adminRepo.GetAdmin(admin.StudentNumber); user != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}
	if err := service.adminRepo.CreateAdmin(&admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}

func (service *Service) GetAdmin(id string) (*model.Admin, error) {
	admin, err := service.adminRepo.GetAdmin(id)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (service *Service) GetAdminById(c *gin.Context) {
	type Request struct {
		Id  string `json:"id"`
		Pwd string `json:"pwd"`
	}
	var input Request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	admin, err := service.adminRepo.GetAdmin(input.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if input.Pwd != admin.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password incorrect"})
		return
	}
	admin.Password = "secret"
	c.JSON(http.StatusOK, admin)
}
