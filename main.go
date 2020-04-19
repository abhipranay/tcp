package main

import (
	"log"

	"abhi.com/tcp/protocol"
	"abhi.com/tcp/tcpserver"

	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	s := &tcpserver.Server{
		Proto:   "tcp",
		Addr:    "localhost:7878",
		Handler: protocol.ProtobufMessageHandler,
	}
	err := s.ListenAndGo().Error()
	if len(err) > 0 {
		panic("Failed to start server: " + err)
	}
	select {}
}
