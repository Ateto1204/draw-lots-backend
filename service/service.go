package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (service *Service) Login(c *gin.Context) {
	type Login struct {
		Identity string `json:"identity"`
		Id       string `json:"id"`
		Pwd      string `json:"pwd"`
	}
	var request Login
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	switch request.Identity {
	case "admin":
		response, err := service.GetAdmin(request.Id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if request.Pwd != response.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "password incorrect"})
			return
		}
		c.JSON(http.StatusOK, response)
		return
	case "senior":
		response, err := service.GetSenior(request.Id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if request.Pwd != response.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "password incorrect"})
			return
		}
		c.JSON(http.StatusOK, response)
		return
	case "junior":
		response, err := service.GetJunior(request.Id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if request.Pwd != response.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "password incorrect"})
			return
		}
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid identity"})
}

func (service *Service) CreateConnect(c *gin.Context) {
	type Connect struct {
		ParentId  string `json:"parent_id"`
		ParentPwd string `json:"parent_pwd"`
		ChildId   string `json:"child_id"`
		ChildPwd  string `json:"child_pwd"`
	}
	var input Connect
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddChildIdToSenior(input.ParentId, input.ChildId, input.ParentPwd); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddParentIdToJunior(input.ParentId, input.ChildId, input.ChildPwd); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (service *Service) ClearConnection(c *gin.Context) {
	type Request struct {
		Id  string `json:"id"`
		Pwd string `json:"pwd"`
	}
	var input Request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
