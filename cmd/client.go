package main

import (
	"fmt"

	"abhi.com/tcp/client"
	"abhi.com/tcp/message"
)

func main() {
	e, p := "abhipranay.chauhan@gmail.com", "asdf@123"
	msg := &message.Login{
		Email:    &e,
		Password: &p,
	}

	c := &client.Client{
		Proto: "tcp",
		SAddr: "0.0.0.0:7878",
	}
	if err := c.Connect(); err != nil {
		panic("Client error: " + err.Error())
	}
	i := 0
	for i < 100 {
		i++
		c.Send(msg)
	}
	fmt.Print(i)
}
