//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

// Если летит ошибка(или еще какие-либо) .../.../main.go:4:7: cannot initialize 1 variables with 2 values
// Тогда нужно просто закомментировать код инициализации инжектора в данном случае
// в файле main я все закомментировал и wire смог собрать wire_gen

// InitializeEvent инжектор
func InitializeEvent(phrase string) (Event, error) {
	// Провайдеры
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
