package gateway

import (
	"github.com/nats-io/nats.go"
)

type MailGw interface {
	SendUpdate(filename string) error
}

type MailGateway struct {
	natsConn *nats.Conn
}

func NewMailGateway(natsConn *nats.Conn) *MailGateway {
	return &MailGateway{
		natsConn: natsConn,
	}
}

func (gw *MailGateway) SendUpdate(filename string) error {
	err := gw.natsConn.Publish("thourus.service.mailman.document_change",
		[]byte("{\n        \"document\": \"Contract\",\n        \"initiator\": \"Maria Petrova\",\n        \"reviewers_emails\": [\"maria_petrova@gmail.com\", \"princess_carolin@gmail.com\"],\n        \"project\": \"Thourus\",\n        \"date\": \"2022-05-01T15:04:05Z\",\n        \"comment\": \"\"\n}"))
	if err != nil {
		return err
	}
	return nil
}
