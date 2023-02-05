package support

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yafgo/framework/contracts/queue"
	"github.com/yafgo/framework/support/file"
	testingfile "github.com/yafgo/framework/testing/file"
)

type Test struct {
}

// Signature The name and signature of the job.
func (receiver *Test) Signature() string {
	return "test"
}

// Handle Execute the job.
func (receiver *Test) Handle(args ...any) error {
	file.Create("test.txt", []byte(args[0].(string)))

	return nil
}

func TestDispatchSync(t *testing.T) {
	task := &Task{
		Job: &Test{},
		Args: []queue.Arg{
			{Type: "uint64", Value: "test"},
		},
	}

	err := task.DispatchSync()
	assert.Nil(t, err)
	assert.True(t, file.Exists("test.txt"))
	assert.True(t, testingfile.GetLineNum("test.txt") == 1)
	res := file.Remove("test.txt")
	assert.True(t, res)
}
