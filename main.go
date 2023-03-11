package main

import (
	"log"

	"github.com/FUnigrad/funiverse-workspace-service/config"
	"github.com/FUnigrad/funiverse-workspace-service/goclient"
	"github.com/FUnigrad/funiverse-workspace-service/model"
)

func main() {

	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Cannot load config: ", err)
	}

	client, err := goclient.NewClient(config)
	if err != nil {
		log.Fatalln("Cannot init K8s Client:", err)
	}

	// fmt.Println(client)

	workspace := model.Workspace{
		Code: "fudn",
	}

	err = client.CreateWorkspace(workspace)
	if err != nil {
		log.Fatalln(err)
	}
	// server := handler.NewServer(client)

	// err = server.Start(config)
	// if err != nil {
	// 	log.Fatalln("Cannot start server:", err)
	// }
}
