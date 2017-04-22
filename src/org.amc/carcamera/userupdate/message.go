package userupdate

import ()

type MessageService interface {
	Init() error
	Error(message string)
	Started() //message that the monitored program has started
	Stopped() //message that the monitored program has stopped
	Close()
}

type Message struct {
	services []MessageService
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

func (m *Message) AddService(service MessageService) {
	m.services = append(m.services, service)
}

func (m *Message) Close() {
	for _, service := range m.services {
		service.Close()
	}
}
