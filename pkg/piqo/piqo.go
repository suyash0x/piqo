package piqo

import (
	"github.com/suyash0x/piqo/pkg/server"
)

type serverModule interface {
	StartServer()
}

type Piqo struct {
	serverModule
}

func newPiqo() *Piqo {
	return &Piqo{
		serverModule: server.NewServer(),
	}
}

func New() *Piqo {
	return newPiqo()

}
