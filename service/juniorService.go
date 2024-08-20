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

func (service *Service) GetAllJuniors(c *gin.Context) {
	juniors, err := service.juniorRepo.GetAllJuniors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, juniors)
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

func (service *Service) AddParentIdToJunior(c *gin.Context) {
	id := c.Param("id")

	junior, err := service.juniorRepo.GetJunior(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	type SeniorId struct {
		Id string `json:"id"`
	}
	var seniorId SeniorId
	if err := c.ShouldBindJSON(&seniorId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	junior.ParentId = seniorId.Id
	if err := service.juniorRepo.UpdateParentId(junior); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		return
	}
	c.JSON(http.StatusOK, junior)
}

func (service *Service) AddLineIdToJunior(c *gin.Context) {
	id := c.Param("id")

	junior, err := service.juniorRepo.GetJunior(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error:": err.Error()})
		return
	}

	type LineId struct {
		Id string `json:"id"`
	}
	var lineId LineId
	if err := c.ShouldBindJSON(&lineId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	junior.LineId = lineId.Id
	if err := service.juniorRepo.UpdateLineId(junior); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		return
	}
	c.JSON(http.StatusOK, junior)
}
