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

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, "NTOUCSE113")
}

func Api(c *gin.Context) {
	type Option struct {
		Admin  string
		Senior string
		Junior string
	}
	c.JSON(http.StatusOK, &Option{})
}

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

func (service *Service) GetAdmin(c *gin.Context) {
	id := c.Param("id")
	admin, err := service.adminRepo.GetAdmin(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
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

func (service *Service) GetJunior(c *gin.Context) {
	id := c.Param("id")
	junior, err := service.juniorRepo.GetJunior(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, junior)
}
