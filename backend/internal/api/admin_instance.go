package api

import (
	"errors"
	"net/http"

	"swjtu-ctf-oj/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (a *ChallengeAPI) AdminListInstances(c *gin.Context) {
	instances, err := a.service.ListAdminContainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to fetch instances: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(instances))
}

func (a *ChallengeAPI) AdminDestroyInstance(c *gin.Context) {
	instanceID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	instance, err := a.service.AdminDestroyContainer(instanceID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, model.ErrorResponse(404, "Instance not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to destroy instance: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(instance))
}
