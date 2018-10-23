package userupdate

import (
	"log"
)

var (
	consolePrefix                = "Web App"
	consoleSerivceStartMessage   = "Console logging started"
	consoleSerivceStoppedMessage = "Console logging stopped"
)

type ConsoleService struct{}

func (l *ConsoleService) Init() error {
	return nil
}

func (c ConsoleService) Error(message string) {
	log.Printf("%s:%s", consolePrefix, message)
}

func (c ConsoleService) Started() {
	log.Println(consoleSerivceStartMessage)
}

func (c ConsoleService) Stopped() {
	log.Println(consoleSerivceStoppedMessage)
}

func (c ConsoleService) Close() {
	// no resources to tidy up
}
