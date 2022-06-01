package gateway

import (
	"github.com/nats-io/nats.go"
)

type MailGw interface {
}

type MailGateway struct {
	natsConn *nats.Conn
}

func NewMailGateway(natsConn *nats.Conn) *MailGateway {
	return &MailGateway{
		natsConn: natsConn,
	}
}
