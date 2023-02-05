package console

import (
	"errors"
	"os"
	"strings"

	"github.com/gookit/color"

	"github.com/yafgo/framework/contracts/console"
	"github.com/yafgo/framework/contracts/console/command"
	"github.com/yafgo/framework/support/file"
	"github.com/yafgo/framework/support/str"
)

type JobMakeCommand struct {
}

// Signature The name and signature of the console command.
func (receiver *JobMakeCommand) Signature() string {
	return "make:job"
}

// Description The console command description.
func (receiver *JobMakeCommand) Description() string {
	return "Create a new job class"
}

// Extend The console command extend.
func (receiver *JobMakeCommand) Extend() command.Extend {
	return command.Extend{
		Category: "make",
	}
}

// Handle Execute the console command.
func (receiver *JobMakeCommand) Handle(ctx console.Context) error {
	name := ctx.Argument(0)
	if name == "" {
		return errors.New("no enough arguments (missing: name) ")
	}

	file.Create(receiver.getPath(name), []byte(receiver.populateStub(receiver.getStub(), name)))
	color.Greenln("Job created successfully")

	return nil
}

func (receiver *JobMakeCommand) getStub() string {
	return JobStubs{}.Job()
}

// populateStub Populate the place-holders in the command stub.
func (receiver *JobMakeCommand) populateStub(stub string, name string) string {
	stub = strings.ReplaceAll(stub, "DummyJob", str.Camel(name))
	stub = strings.ReplaceAll(stub, "DummyName", str.Camel(name))

	return stub
}

// getPath Get the full path to the command.
func (receiver *JobMakeCommand) getPath(name string) string {
	pwd, _ := os.Getwd()

	return pwd + "/app/jobs/" + str.Snake(name) + ".go"
}
