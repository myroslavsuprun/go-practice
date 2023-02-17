package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, greeting Greeter) string {
	return fmt.Sprintf("I can speak %v: %v", greeting.LanguageName(), greeting.Greet(name))
}

type Italian struct{}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

func (i Italian) LanguageName() string {
	return "Italian"
}

type Portuguese struct{}

func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}
