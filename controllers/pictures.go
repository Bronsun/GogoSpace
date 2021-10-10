package controllers

import (
	"errors"
	"net/http"

	"github.com/Bronsun/GogoSpace/models"
	"github.com/Bronsun/GogoSpace/request"
	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidRequest    = errors.New("Invalid request")
	ErrFailedToGetImages = errors.New("Failed to get images")
)

type PicturesController struct{}

func (p PicturesController) GetPictures(ctx *gin.Context) {
	var req models.Date
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidRequest})
		return
	}

	if err := req.ValidateDate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	urls, err := request.GetImagesFromRequest(*req.From, *req.To)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": ErrFailedToGetImages})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"urls": urls,
	})

}
