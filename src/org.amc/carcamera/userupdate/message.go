package userupdate 

import (

)

type messageService interface {
	Init() error
	Error(message string)
	Started()
	Stopped()
	
}

type Message struct {
	services []messageService
}

func (m Message) Error(message string) {
	for _, service := range m.services {
		service.Error(message)
	}
}

/*
 * Initialising the Messaging services
 */
func (m Message) Init() error {
	for _, service := range m.services {
		err := service.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m Message) Started() {
	for _, service := range m.services {
		service.Started()
	}
}

func (m Message) Stopped() {
	for _, service := range m.services {
		service.Stopped()
	}
}


func (m *Message) AddService(service messageService) {
	m.services = append(m.services, service)
} 