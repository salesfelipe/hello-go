package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, greeter Greeter) string {
	return greeter.Greet(name)
}

type Italian struct{}

func (ig Italian) LanguageName() string {
	return "Italian"
}

func (ig Italian) Greet(name string) string {
	return fmt.Sprintf("I can speak Italian: Ciao %s!", name)
}

type Portuguese struct{}

func (ig Portuguese) LanguageName() string {
	return "Portuguese"
}

func (ig Portuguese) Greet(name string) string {
	return fmt.Sprintf("I can speak Portuguese: Olá %s!", name)
}
