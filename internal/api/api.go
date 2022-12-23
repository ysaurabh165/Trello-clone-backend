package api

import (
	"fmt"
	"os"
	"github.com/jacobstr/confer"
	"github.com/ysaurabh165/Trello-clone-backend/internal/router"
)

var configuration *confer.Config
var mode string
var port string

func Run(configPath string) {
	if len(os.Args) > 1 {
		mode = os.Args[1]
	}
	if len(os.Args) > 2 {
		port = os.Args[2]
	}
	fmt.Println("mode: "+mode)
	fmt.Println("port: "+port)
	if configPath == "" {
		configPath = "C:/Users/ysaur/OneDrive/Desktop/Trello-clone-backend/data/config.yaml"
	}

	configuration = confer.NewConfig()
	loadDefaultConfig(configPath, mode)
	web := router.Setup(configuration)
	_ = web.Run("localhost:"+port)
}

func loadDefaultConfig(configPath string, runMode string) {

	if runMode == "prod" {
		mode = runMode
	} else {
		mode = "dev"
	}

	err := configuration.ReadPaths(configPath)
	if err == nil {
		fmt.Println("Configuration loaded :" + mode)
	} else {
		fmt.Println(err.Error())
		fmt.Println("No configuration file found")
	}

}