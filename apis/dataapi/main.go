package main

import (
	"antinolabsassignment/apis/dataapi/controllers"
	"antinolabsassignment/pkg/common/router"
	"github.com/spf13/viper"
	"log"
)

func main() {
	vp := viper.New()
	port := vp.GetString("DATA_API")
	if len(port) == 0 {
		port = ":8080"
	}

	httpRouter := router.NewMuxRouter()
	controllers.RegisterControllers(httpRouter)
	err := httpRouter.Serve("DATA_API", port)
	if err != nil {
		log.Fatalf("data api failed to start")
	}
}
