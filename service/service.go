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
		Id     string `json:"id"`
		Pwd    string `json:"pwd"`
		Parent string `json:"parent"`
		Child  string `json:"child"`
	}
	var input Connect
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authorized by admin
	admin, err := service.adminRepo.GetAdmin(input.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if input.Pwd != admin.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password incorrect"})
		return
	}

	// Create connection
	if err := service.AddChildIdToSenior(input.Parent, input.Child); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddParentIdToJunior(input.Parent, input.Child); err != nil {
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
	// Authorized by admin
	admin, err := service.adminRepo.GetAdmin(input.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if input.Pwd != admin.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password incorrect"})
		return
	}

	// Clear seniors
	seniors, err := service.seniorRepo.GetAllSeniors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	seniorList := *seniors
	for i := range seniorList {
		seniorList[i].ChildrenId = model.StringArray{}
	}

	for _, senior := range seniorList {
		err := service.seniorRepo.UpdateChildId(&senior)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Clear juniors
	juniors, err := service.juniorRepo.GetAllJuniors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	juniorList := *juniors
	for i := range juniorList {
		juniorList[i].ParentId = ""
	}

	for _, junior := range juniorList {
		err := service.juniorRepo.UpdateParentId(&junior)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
