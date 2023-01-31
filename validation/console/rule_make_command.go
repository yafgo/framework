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

type RuleMakeCommand struct {
}

// Signature The name and signature of the console command.
func (receiver *RuleMakeCommand) Signature() string {
	return "make:rule"
}

// Description The console command description.
func (receiver *RuleMakeCommand) Description() string {
	return "Create a new rule class"
}

// Extend The console command extend.
func (receiver *RuleMakeCommand) Extend() command.Extend {
	return command.Extend{
		Category: "make",
	}
}

// Handle Execute the console command.
func (receiver *RuleMakeCommand) Handle(ctx console.Context) error {
	name := ctx.Argument(0)
	if name == "" {
		return errors.New("no enough arguments (missing: name) ")
	}

	file.Create(receiver.getPath(name), []byte(receiver.populateStub(receiver.getStub(), name)))
	color.Greenln("Rule created successfully")

	return nil
}

func (receiver *RuleMakeCommand) getStub() string {
	return Stubs{}.Request()
}

// populateStub Populate the place-holders in the command stub.
func (receiver *RuleMakeCommand) populateStub(stub string, name string) string {
	stub = strings.ReplaceAll(stub, "DummyRule", str.Camel(name))
	stub = strings.ReplaceAll(stub, "DummyName", str.Snake(name))

	return stub
}

// getPath Get the full path to the command.
func (receiver *RuleMakeCommand) getPath(name string) string {
	pwd, _ := os.Getwd()

	return pwd + "/app/rules/" + str.Snake(name) + ".go"
}
