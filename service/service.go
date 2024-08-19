package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/we-we-Web/draw-lots-backend/model"
	"github.com/we-we-Web/draw-lots-backend/repository"
)

type Service struct {
	adminRepo  *repository.AdminRepository
	seniorRepo *repository.SeniorRepository
	juniorRepo *repository.JuniorRepository
}

func NewService(adminRepo *repository.AdminRepository,
	seniorRepo *repository.SeniorRepository,
	juniorRepo *repository.JuniorRepository) *Service {

	return &Service{
		adminRepo:  adminRepo,
		seniorRepo: seniorRepo,
		juniorRepo: juniorRepo,
	}
}

func (service *Service) CreateAdmin(c *gin.Context) {
	var admin model.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.adminRepo.CreateAdmin(&admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}

func (service *Service) CreateSenior(c *gin.Context) {

}

func (service *Service) CreateJunior(c *gin.Context) {

}

func (service *Service) GetAdmin(c *gin.Context) {
	id := c.Param("id")
	admin := service.adminRepo.GetAdmin(id)
	c.JSON(http.StatusOK, admin)
}

func (service *Service) GetSenior(c *gin.Context) {
	id := c.Param("id")
	senior := service.seniorRepo.GetSenior(id)
	c.JSON(http.StatusOK, senior)
}

func (service *Service) GetJunior(c *gin.Context) {
	id := c.Param("id")
	junior := service.juniorRepo.GetJunior(id)
	c.JSON(http.StatusOK, junior)
}
