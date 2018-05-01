package command

import (
	"bytes"
	"fmt"
	proto "github.com/chremoas/chremoas/proto"
	esisvc "github.com/chremoas/esi-srv/proto"
	"golang.org/x/net/context"
	"strings"
)

type ClientFactory interface {
	NewSearchServiceClient() esisvc.SearchService
}

var clientFactory ClientFactory

type Command struct {
	//Store anything you need the Help or Exec functions to have access to here
	name    string
	factory ClientFactory
}

func (c *Command) Help(ctx context.Context, req *proto.HelpRequest, rsp *proto.HelpResponse) error {
	rsp.Usage = c.name
	rsp.Description = "Lookup anything in ESI"
	return nil
}

func (c *Command) Exec(ctx context.Context, req *proto.ExecRequest, rsp *proto.ExecResponse) error {
	if len(req.Args) == 1 {
		rsp.Result = []byte("Need an argument")
		return nil
	}

	if req.Args[1] == "help" {
		var buffer bytes.Buffer

		buffer.WriteString("Usage: !lookup <category> <search string>\n")
		buffer.WriteString("\tcategory: The entity category to look for\n")
		buffer.WriteString("\tsearch string: The entity to search for (needs to be at least three characters\n")

		rsp.Result = []byte(fmt.Sprintf("```%s```", buffer.String()))
		return nil
	}

	client := clientFactory.NewSearchServiceClient()

	//category := req.Args[1]
	//fmt.Printf("Category: %s\n", category)
	searchString := strings.Join(req.Args[2:], " ")
	//fmt.Printf("searchString: %s\n", searchString)
	sr := esisvc.SearchRequest{SearchString: searchString}
	response, err := client.Search(ctx, &sr)

	if err != nil {
		fmt.Printf("Something be wrong, yo! %+v\n", err)
	} else {
		fmt.Printf("response: %+v\n", response)
	}

	rsp.Result = []byte(fmt.Sprintf("```%+v```", response))
	return nil
}

func NewCommand(name string, factory ClientFactory) *Command {
	clientFactory = factory
	newCommand := Command{name: name, factory: factory}
	return &newCommand
}
