package main

import (
	"errors"
	"fmt"
	"github.com/google/wire"
	"time"
)

var SuperSet = wire.NewSet(NewEvent, NewGreeter, NewMessage)

func NewMessage(phrase string) Message {
	return Message(phrase)
}

type Message string

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

type Greeter struct {
	Message Message
	Grumpy  bool
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		if g.Grumpy {
			return Event{}, errors.New("could not create event: event greeter is grumpy")
		}
	}
	return Event{Greeter: g}, nil
}

type Event struct {
	Greeter Greeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
