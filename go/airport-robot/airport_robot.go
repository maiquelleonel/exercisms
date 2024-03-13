package airportrobot

import (
	"fmt"
)

type Italian struct{}

type Portuguese struct{}

type Greeter interface {
	LanguageName() string
	Greet(visitorName string) string
}

func (l Italian) LanguageName() string {
	return "Italian"
}

func (l Italian) Greet(visitorName string) string {
	return fmt.Sprintf("Ciao %s!", visitorName)
}

func (l Portuguese) LanguageName() string {
	return "Portuguese"
}

func (l Portuguese) Greet(visitorName string) string {
	return fmt.Sprintf("Olá %s!", visitorName)
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}
