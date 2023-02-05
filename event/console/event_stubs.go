package console

type EventStubs struct {
}

func (receiver EventStubs) Event() string {
	return `package events

import "github.com/yafgo/framework/contracts/events"

type DummyEvent struct {
}

func (receiver *DummyEvent) Handle(args []events.Arg) ([]events.Arg, error) {
	return args, nil
}
`
}
