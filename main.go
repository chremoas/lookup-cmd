package main

import (
	"fmt"

	proto "github.com/chremoas/chremoas/proto"
	esisvc "github.com/chremoas/esi-srv/proto"
	"github.com/chremoas/services-common/config"
	chremoasPrometheus "github.com/chremoas/services-common/prometheus"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"go.uber.org/zap"

	"github.com/chremoas/lookup-cmd/command"
)

var (
	Version = "SET ME YOU KNOB"
	service micro.Service
	name    = "lookup"
	logger  *zap.Logger
)

func main() {
	var err error

	// TODO pick stuff up from the config
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("Initialized logger")

	go chremoasPrometheus.PrometheusExporter(logger)

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
