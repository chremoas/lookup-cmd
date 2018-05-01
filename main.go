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

var Version = "SET ME YOU KNOB"
var service micro.Service
var name = "lookup"

func main() {
	service = config.NewService(Version, "cmd", name, initialize)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

// This function is a callback from the config.NewService function.  Read those docs
func initialize(config *config.Configuration) error {
	clientFactory := clientFactory{name: config.LookupService("srv", "esi"), client: service.Client()}

	proto.RegisterCommandHandler(service.Server(),
		command.NewCommand(name,
			&clientFactory,
		),
	)

	return nil
}

type clientFactory struct {
	name   string
	client client.Client
}

func (c clientFactory) NewSearchServiceClient() esisvc.SearchService {
	return esisvc.NewSearchService(c.name, c.client)
}
