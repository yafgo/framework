package console

import (
	"errors"
	"os"
	"strings"

	"github.com/yafgo/framework/contracts/console"
	"github.com/yafgo/framework/contracts/console/command"
	"github.com/yafgo/framework/support/file"
	"github.com/yafgo/framework/support/str"

	"github.com/gookit/color"
)

type RequestMakeCommand struct {
}

// Signature The name and signature of the console command.
func (receiver *RequestMakeCommand) Signature() string {
	return "make:request"
}

// Description The console command description.
func (receiver *RequestMakeCommand) Description() string {
	return "Create a new request class"
}

// Extend The console command extend.
func (receiver *RequestMakeCommand) Extend() command.Extend {
	return command.Extend{
		Category: "make",
	}
}

// Handle Execute the console command.
func (receiver *RequestMakeCommand) Handle(ctx console.Context) error {
	name := ctx.Argument(0)
	if name == "" {
		return errors.New("no enough arguments (missing: name) ")
	}

	file.Create(receiver.getPath(name), []byte(receiver.populateStub(receiver.getStub(), name)))
	color.Greenln("Request created successfully")

	return nil
}

func (receiver *RequestMakeCommand) getStub() string {
	return Stubs{}.Request()
}

// populateStub Populate the place-holders in the command stub.
func (receiver *RequestMakeCommand) populateStub(stub string, name string) string {
	stub = strings.ReplaceAll(stub, "DummyRequest", str.Camel(name))
	stub = strings.ReplaceAll(stub, "DummyField", "Name string `form:\"name\" json:\"name\"`")

	return stub
}

// getPath Get the full path to the command.
func (receiver *RequestMakeCommand) getPath(name string) string {
	pwd, _ := os.Getwd()

	return pwd + "/app/http/requests/" + str.Snake(name) + ".go"
}
