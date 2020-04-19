package tcpserver

import (
	"testing"

	"abhi.com/tcp/client"
	"abhi.com/tcp/message"
)

func BenchmarkServer(t *testing.B) {
	sendMessage()
}

func sendMessage() {
	e, p := "abhipranay.chauhan@gmail.com", "asdf@123"
	msg := &message.Login{
		Email:    e,
		Password: p,
	}

	c := &client.Client{
		Proto: "tcp",
		SAddr: "0.0.0.0:7878",
	}
	if err := c.Connect(); err != nil {
		panic("Client error: " + err.Error())
	}
	i := 0
	for i < 1000 {
		i++
		c.Send(msg)
	}
}
