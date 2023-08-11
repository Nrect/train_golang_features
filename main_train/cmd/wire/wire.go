//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// InitializeEvent инжектор
func InitializeEvent() Event {
	// Провайдеры
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
