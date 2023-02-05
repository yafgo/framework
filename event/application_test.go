package event

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/yafgo/framework/config"
	"github.com/yafgo/framework/contracts/event"
	eventcontract "github.com/yafgo/framework/contracts/event"
	"github.com/yafgo/framework/facades"
	"github.com/yafgo/framework/queue"
	testingdocker "github.com/yafgo/framework/testing/docker"
)

var (
	testSyncListener        = 0
	testAsyncListener       = 0
	testCancelListener      = 0
	testCancelAfterListener = 0
)

type EventTestSuite struct {
	suite.Suite
}

func TestEventTestSuite(t *testing.T) {
	redisPool, redisResource, err := testingdocker.Redis()
	if err != nil {
		log.Fatalf("Get redis error: %s", err)
	}

	initConfig(redisResource.GetPort("6379/tcp"))
	facades.Queue = queue.NewApplication()
	facades.Event = NewApplication()

	suite.Run(t, new(EventTestSuite))

	if err := redisPool.Purge(redisResource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func (s *EventTestSuite) SetupTest() {

}

func (s *EventTestSuite) TestEvent() {
	facades.Event.Register(map[event.Event][]event.Listener{
		&TestEvent{}: {
			&TestSyncListener{},
			&TestAsyncListener{},
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		s.Nil(facades.Queue.Worker(nil).Run())

		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	s.Nil(facades.Event.Job(&TestEvent{}, []eventcontract.Arg{
		{Type: "string", Value: "Yafgo"},
		{Type: "int", Value: 1},
	}).Dispatch())
	time.Sleep(1 * time.Second)
	s.Equal(1, testSyncListener)
	s.Equal(1, testAsyncListener)
}

func (s *EventTestSuite) TestCancelEvent() {
	facades.Event.Register(map[event.Event][]event.Listener{
		&TestCancelEvent{}: {
			&TestCancelListener{},
			&TestCancelAfterListener{},
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		s.Nil(facades.Queue.Worker(nil).Run())

		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	s.EqualError(facades.Event.Job(&TestCancelEvent{}, []eventcontract.Arg{
		{Type: "string", Value: "Yafgo"},
		{Type: "int", Value: 1},
	}).Dispatch(), "cancel")
	time.Sleep(1 * time.Second)
	s.Equal(1, testCancelListener)
	s.Equal(0, testCancelAfterListener)
}

func initConfig(redisPort string) {
	application := config.NewApplication("../.env")
	application.Add("app", map[string]any{
		"name": "yafgo",
	})
	application.Add("queue", map[string]any{
		"default": "redis",
		"connections": map[string]any{
			"sync": map[string]any{
				"driver": "sync",
			},
			"redis": map[string]any{
				"driver":     "redis",
				"connection": "default",
				"queue":      "default",
			},
		},
	})
	application.Add("database", map[string]any{
		"redis": map[string]any{
			"default": map[string]any{
				"host":     "localhost",
				"password": "",
				"port":     redisPort,
				"database": 0,
			},
		},
	})

	facades.Config = application
}

type TestEvent struct {
}

func (receiver *TestEvent) Handle(args []event.Arg) ([]event.Arg, error) {
	return args, nil
}

type TestCancelEvent struct {
}

func (receiver *TestCancelEvent) Handle(args []event.Arg) ([]event.Arg, error) {
	return args, nil
}

type TestAsyncListener struct {
}

func (receiver *TestAsyncListener) Signature() string {
	return "test_async_listener"
}

func (receiver *TestAsyncListener) Queue(args ...any) event.Queue {
	return event.Queue{
		Enable:     true,
		Connection: "",
		Queue:      "",
	}
}

func (receiver *TestAsyncListener) Handle(args ...any) error {
	testAsyncListener++

	return nil
}

type TestSyncListener struct {
}

func (receiver *TestSyncListener) Signature() string {
	return "test_sync_listener"
}

func (receiver *TestSyncListener) Queue(args ...any) event.Queue {
	return event.Queue{
		Enable:     false,
		Connection: "",
		Queue:      "",
	}
}

func (receiver *TestSyncListener) Handle(args ...any) error {
	testSyncListener++

	return nil
}

type TestCancelListener struct {
}

func (receiver *TestCancelListener) Signature() string {
	return "test_cancel_listener"
}

func (receiver *TestCancelListener) Queue(args ...any) event.Queue {
	return event.Queue{
		Enable:     false,
		Connection: "",
		Queue:      "",
	}
}

func (receiver *TestCancelListener) Handle(args ...any) error {
	testCancelListener++

	return errors.New("cancel")
}

type TestCancelAfterListener struct {
}

func (receiver *TestCancelAfterListener) Signature() string {
	return "test_cancel_after_listener"
}

func (receiver *TestCancelAfterListener) Queue(args ...any) event.Queue {
	return event.Queue{
		Enable:     false,
		Connection: "",
		Queue:      "",
	}
}

func (receiver *TestCancelAfterListener) Handle(args ...any) error {
	testCancelAfterListener++

	return nil
}
