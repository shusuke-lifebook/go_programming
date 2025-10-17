package controllers

import (
	"go-api-arch-mvc-template/api"
	"go-api-arch-mvc-template/app/models"
	"go-api-arch-mvc-template/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct{}

func (a *AlbumHandler) CreateAlbum(c *gin.Context) {
	var requestBody api.CreateAlbumJSONRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Warn(err.Error())
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
		return
	}

	createdAlbum, err := models.CreateAlbum(
		requestBody.Title,
		requestBody.ReleaseDate.Time,
		string(requestBody.Category.Name))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdAlbum)
}
