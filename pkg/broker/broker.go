package broker

import (
	"log"
	"net"
)

type Broker struct{}

func (b *Broker) AddTopic(conn net.Conn) {
	log.Println("Add topic request received")
}
func (b *Broker) PostMessage(conn net.Conn) {
	log.Println("Post request received")
}
func (b *Broker) GetMessage(conn net.Conn) {
	log.Println("Get request received")
}

func NewBroker() *Broker {
	return &Broker{}
}
