package nats

import "github.com/nats-io/nats.go"

func (nc *Client) Publish(subject string, data []byte) error {
	return nc.connection.Publish(subject, data)
}

func (nc *Client) Subscribe(subject string, handler func(msg []byte)) error {
	_, err := nc.connection.Subscribe(subject, func(msg *nats.Msg) {
		handler(msg.Data)
	})
	return err
}

func (nc *Client) Close() {
	nc.connection.Close()
}

func (nc *Client) CheckConnection() string {
	if nc.connection.IsConnected() {
		return "connected"
	}
	return "disconnected"
}
