package main

import (
	"fmt"
	"github.com/chremoas/lookup-cmd/command"
	esisvc "github.com/chremoas/esi-srv/proto"
	proto "github.com/chremoas/chremoas/proto"
	"github.com/chremoas/services-common/config"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

var Version = "1.0.0"
var service micro.Service

func main() {
	service = config.NewService(Version, "lookup", initialize)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

// This function is a callback from the config.NewService function.  Read those docs
func initialize(config *config.Configuration) error {
	authSvcName := config.Bot.AuthSrvNamespace + "." + config.ServiceNames.AuthSrv
	clientFactory := clientFactory{name: authSvcName, client: service.Client()}

	proto.RegisterCommandHandler(service.Server(),
		command.NewCommand(config.Name,
			&clientFactory,
		),
	)

	return nil
}

type clientFactory struct {
	name   string
	client client.Client
}

func (c clientFactory) NewSearchServiceClient() esisvc.SearchServiceClient {
	return esisvc.NewSearchServiceClient(c.name, c.client)
}
