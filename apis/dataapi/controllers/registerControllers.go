package controllers

import (
	"antinolabsassignment/pkg/common/router"
)

//region public functions

// RegisterControllers - it registers all the controllers for this API
// Param router - it points to a router
func RegisterControllers(router *router.MuxRouter) {
	controllers := Controllers{}
	router.POST("/blog/post", controllers.Post)
	router.GET("/blog/get", controllers.GetById)
	router.DELETE("/blog/delete", controllers.DeleteById)
	router.PUT("/blog/update", controllers.UpdateById)
}

//endregion
