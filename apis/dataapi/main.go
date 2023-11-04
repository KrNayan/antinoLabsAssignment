package main

import (
	"antinolabsassignment/apis/dataapi/controllers"
	"antinolabsassignment/pkg/common/router"
	"antinolabsassignment/pkg/common/utilities/viper"
	"log"
)

func main() {
	vp, err := viper.NewViper()
	if err != nil {
		log.Fatalf("data api failed to start")
	}
	port := vp.GetString("DATA_API_PORT")
	if len(port) == 0 {
		port = ":8080"
	}

	httpRouter := router.NewMuxRouter()
	controllers.RegisterControllers(httpRouter)
	err = httpRouter.Serve("DATA_API", port)
	if err != nil {
		log.Fatalf("data api failed to start")
	}
}
