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
	NewSearchServiceClient() esisvc.SearchServiceClient
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

	client := clientFactory.NewSearchServiceClient()

	category := req.Args[1]
	fmt.Printf("Category: %s\n", category)
	searchString := strings.Join(req.Args[2:], " ")
	fmt.Printf("searchString: %s\n", searchString)
	sr := esisvc.SearchRequest{SearchString: searchString}
	response, err := client.Search(ctx, &sr)

	if err != nil {
		fmt.Printf("Something be wrong, yo! %+v\n", err)
	} else {
		fmt.Printf("response: %+v\n", response)
	}

	//rsp.Result = []byte(response)
	return nil
}

func help(ctx context.Context, req *proto.ExecRequest) string {
	var buffer bytes.Buffer

	buffer.WriteString("Usage: !lookup <type> <arguments>\n")
	buffer.WriteString("\thelp: This text\n")

	return fmt.Sprintf("```%s```", buffer.String())
}

//func addRole(ctx context.Context, req *proto.ExecRequest) string {
//	var buffer bytes.Buffer
//	client := clientFactory.NewEntityAdminClient()
//	roleName := req.Args[2]
//	chatServiceGroup := strings.Join(req.Args[2:], " ")
//
//	if len(chatServiceGroup) > 0 && chatServiceGroup[0] == '"' {
//		chatServiceGroup = chatServiceGroup[1:]
//	}
//	if len(chatServiceGroup) > 0 && chatServiceGroup[len(chatServiceGroup)-1] == '"' {
//		chatServiceGroup = chatServiceGroup[:len(chatServiceGroup)-1]
//	}
//
//	output, err := client.RoleUpdate(ctx, &uauthsvc.RoleAdminRequest{
//		Role:      &uauthsvc.Role{RoleName: roleName, ChatServiceGroup: chatServiceGroup},
//		Operation: uauthsvc.EntityOperation_ADD_OR_UPDATE,
//	})
//
//	if err != nil {
//		buffer.WriteString(err.Error())
//	} else {
//		buffer.WriteString(output.String())
//	}
//
//	return fmt.Sprintf("```%s```", buffer.String())
//}

func NewCommand(name string, factory ClientFactory) *Command {
	clientFactory = factory
	newCommand := Command{name: name, factory: factory}
	return &newCommand
}
