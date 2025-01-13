package nats

import (
	"sync"

	"github.com/nats-io/nats.go"
)

var (
	once sync.Once
	nc   *Client
)

type Client struct {
	connection *nats.Conn
}

func NewNATsClient(url string) (*Client, error) {
	var err error
	once.Do(func() {
		conn, err := nats.Connect(url)
		if err != nil {
			return
		}

		nc = &Client{
			connection: conn,
		}
	})
	if err != nil {
		return nil, err
	}
	return nc, nil
}
