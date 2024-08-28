package service

import (
	"math/rand"
	"net/http"
	"time"

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
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}
	if err := service.juniorRepo.CreateJunior(&junior); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, junior)
}

func (service *Service) GetAllJuniors(c *gin.Context) {
	juniors, err := service.juniorRepo.GetAllJuniors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, juniors)
}

func (service *Service) GetJunior(id string) (*model.Junior, error) {
	junior, err := service.juniorRepo.GetJunior(id)
	if err != nil {
		return nil, err
	}
	return junior, nil
}

func (service *Service) AddParentIdToJunior(parentId, childId string) error {
	junior, err := service.juniorRepo.GetJunior(childId)
	if err != nil {
		return err
	}

	junior.ParentId = parentId
	if err := service.juniorRepo.UpdateParentId(junior); err != nil {
		return err
	}
	return nil
}

func (service *Service) SetLineIdToJunior(c *gin.Context) {
	type Request struct {
		Id   string `json:"id"`
		Pwd  string `json:"pwd"`
		Line string `json:"line"`
	}
	var input Request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	junior, err := service.juniorRepo.GetJunior(input.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if input.Pwd != junior.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password incorrect"})
		return
	}

	junior.LineId = input.Line
	if err := service.juniorRepo.UpdateLineId(junior); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, junior)
}

func (service *Service) GetJuniorById(c *gin.Context) {
	id := c.Param("id")
	junior, err := service.juniorRepo.GetJunior(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	junior.Password = "secret"
	c.JSON(http.StatusOK, junior)
}

// MARK: - PickJunior -
func (service *Service) PickJunior(c *gin.Context) {
	juniors, err := service.juniorRepo.GetAllJuniors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	openJuniors := []model.Junior{}
	juniorList := *juniors
	for _, junior := range juniorList {
		if junior.ParentId == "" {
			openJuniors = append(openJuniors, junior)
		}
	}

	if len(openJuniors) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No available juniors found"})
		return
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomJunior := openJuniors[r.Intn(len(openJuniors))]

	randomJunior.Password = "secret"
	c.JSON(http.StatusOK, randomJunior)
}
