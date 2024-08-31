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

// MARK: - NewService -
func NewService(adminRepo *repository.AdminRepository,
	seniorRepo *repository.SeniorRepository,
	juniorRepo *repository.JuniorRepository) *Service {

	return &Service{
		adminRepo:  adminRepo,
		seniorRepo: seniorRepo,
		juniorRepo: juniorRepo,
	}
}

// MARK: - Login -
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

// MARK: - CreateConnect -
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
	errCh := make(chan error, 2)

	go func() {
		errCh <- service.AddChildIdToSenior(input.Parent, input.Child)
	}()
	go func() {
		errCh <- service.AddParentIdToJunior(input.Parent, input.Child)
	}()

	for i := 0; i < 2; i++ {
		if err := <-errCh; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// MARK: - CreateConnectByInvited -
func (service *Service) CreateConnectByInvited(c *gin.Context) {
	type Connect struct {
		Parent string `json:"parent"`
		Pwd    string `json:"pwd"`
		Child  string `json:"child"`
	}
	var input Connect
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authorized by senior
	senior, err := service.seniorRepo.GetSenior(input.Parent)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if input.Pwd != senior.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password incorrect"})
		return
	}

	// Create connection
	errCh := make(chan error, 2)

	go func() {
		errCh <- service.AddChildIdToSenior(input.Parent, input.Child)
	}()
	go func() {
		errCh <- service.AddParentIdToJunior(input.Parent, input.Child)
	}()

	for i := 0; i < 2; i++ {
		if err := <-errCh; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// MARK: - ClearConnection -
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

	errCh := make(chan error, 2)

	// Clear seniors
	go func() {
		seniors, err := service.seniorRepo.GetAllSeniors()
		if err != nil {
			errCh <- err
			return
		}

		seniorList := *seniors
		for i := range seniorList {
			go func(senior *model.Senior) {
				senior.ChildrenId = model.StringArray{}
				if err := service.seniorRepo.UpdateSenior(senior); err != nil {
					errCh <- err
					return
				}
			}(&seniorList[i])
		}

		for _, senior := range seniorList {
			if err := service.seniorRepo.UpdateSenior(&senior); err != nil {
				errCh <- err
				return
			}
		}
		errCh <- nil
	}()

	// Clear juniors
	go func() {
		juniors, err := service.juniorRepo.GetAllJuniors()
		if err != nil {
			errCh <- err
			return
		}

		juniorList := *juniors
		for i := range juniorList {
			go func(junior *model.Junior) {
				junior.ParentId = ""
				if err := service.juniorRepo.UpdateJunior(junior); err != nil {
					errCh <- err
					return
				}
			}(&juniorList[i])
		}

		errCh <- nil
	}()

	for i := 0; i < 2; i++ {
		if err := <-errCh; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
