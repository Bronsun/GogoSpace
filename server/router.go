package server

import (
	"github.com/Bronsun/GogoSpace/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	pictures := new(controllers.PicturesController)

	router.GET("/pictures", pictures.GetPictures)

	return router
}
