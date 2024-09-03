package service

import (
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/we-we-Web/draw-lots-backend/model"
)

// MARK: - CreateSenior -
func (service *Service) CreateSenior(c *gin.Context) {
	var senior model.Senior
	if err := c.ShouldBindJSON(&senior); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user, _ := service.seniorRepo.GetSenior(senior.StudentNumber); user != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}
	if err := service.seniorRepo.CreateSenior(&senior); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, senior)
}

// MARK: - GetAllSeniors -
func (service *Service) GetAllSeniors(c *gin.Context) {
	seniors, err := service.seniorRepo.GetAllSeniors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for i := range *seniors {
		(*seniors)[i].Password = "secret"
	}
	c.JSON(http.StatusOK, seniors)
}

// MARK: - GetSenior -
func (service *Service) GetSenior(id string) (*model.Senior, error) {
	senior, err := service.seniorRepo.GetSenior(id)
	if err != nil {
		return nil, err
	}
	return senior, nil
}

// MARK: - AddChildToSenior -
func (service *Service) AddChildIdToSenior(parentId, childId string) error {
	senior, err := service.seniorRepo.GetSenior(parentId)
	if err != nil {
		return err
	}

	if len(senior.ChildrenId) >= senior.Quota {
		return errors.New("the limit has been reached")
	}
	for _, junior := range senior.ChildrenId {
		if childId == junior {
			return errors.New("the child has already existed")
		}
	}
	senior.ChildrenId = *senior.ChildrenId.Append(childId)
	if err := service.seniorRepo.UpdateSenior(senior); err != nil {
		return err
	}
	return nil
}

// MARK: - AddChildToSeniorByInvite -
func (service *Service) AddChildToSeniorByInvite(parentId, childId string) error {
	senior, err := service.seniorRepo.GetSenior(parentId)
	if err != nil {
		return err
	}
	for _, junior := range senior.ChildrenId {
		if childId == junior {
			return errors.New("the child has already existed")
		}
	}
	senior.ChildrenId = *senior.ChildrenId.Append(childId)
	if err := service.seniorRepo.UpdateSenior(senior); err != nil {
		return err
	}
	return nil
}

// MARK: - GetSeniorById -
func (service *Service) GetSeniorById(c *gin.Context) {
	id := c.Param("id")
	senior, err := service.seniorRepo.GetSenior(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	senior.Password = "secret"
	c.JSON(http.StatusOK, senior)
}

// MARK: - PickSenior -
func (service *Service) PickSenior(c *gin.Context) {
	seniors, err := service.seniorRepo.GetAllSeniors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	openSeniors := []model.Senior{}
	seniorList := *seniors
	for _, senior := range seniorList {
		if len(senior.ChildrenId) < senior.Quota {
			openSeniors = append(openSeniors, senior)
		}
	}

	if len(openSeniors) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No available seniors found"})
		return
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomSenior := openSeniors[r.Intn(len(openSeniors))]

	randomSenior.Password = "secret"
	c.JSON(http.StatusOK, randomSenior)
}

// MARK: - EditSenior -
func (service *Service) EditSenior(c *gin.Context) {
	type Request struct {
		Id        string `json:"id"`
		Pwd       string `json:"pwd"`
		Line      string `json:"line,omitempty"`
		Instagram string `json:"ig,omitempty"`
	}
	var input Request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	senior, err := service.seniorRepo.GetSenior(input.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if input.Pwd != senior.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password incorrect"})
		return
	}

	senior.Line = input.Line
	senior.Instagram = input.Instagram
	if err := service.seniorRepo.UpdateSenior(senior); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, senior)
}
