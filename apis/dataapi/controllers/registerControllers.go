package controllers

import (
	"antinolabsassignment/pkg/common/router"
)

func RegisterControllers(router *router.MuxRouter) {
	controllers := Controllers{}
	router.POST("/blog/post", controllers.Post)
	router.GET("/blog/get/{id}", controllers.GetById)
	router.DELETE("/blog/delete/{id}", controllers.DeleteById)
	router.PUT("/blog/update/{id}", controllers.UpdateById)
}
