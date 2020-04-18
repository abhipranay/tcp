package main

import (
	"abhi.com/tcp/protocol"
	"abhi.com/tcp/tcpserver"
)

func main()  {
	s := &tcpserver.Server{
		Proto: "tcp",
		Addr: "0.0.0.0:7878",
		Handler: protocol.ProtobufMessageHandler,
	}
	err := s.ListenAndGo().Error()
	if len(err) > 0 {
		panic("Failed to start server: " + err)
	}
	select{}
}