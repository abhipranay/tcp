package tcpserver

import (
	"bufio"
	"io"
	"net"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	Proto   string
	Addr    string
	Handler func(rw *bufio.ReadWriter) error
}

func (s *Server) ListenAndGo() error {
	tcpaddr, err := net.ResolveTCPAddr(s.Proto, s.Addr)
	if err != nil {
		log.Panicf("Failed to start server. Error: %v", err)
		return err
	}
	listner, err := net.ListenTCP(s.Proto, tcpaddr)
	if err != nil {
		log.Panic("Failed to listen for tcp connections on address ", s.Addr, " Error: ", err)
		return err
	}
	log.Info("Server started at " + tcpaddr.String())
	for {
		conn, err := listner.AcceptTCP()
		if err != nil {
			log.Error("Failed to accept connection ", conn, " due to error ", err)
			continue
		}
		log.Info("Client ", conn.RemoteAddr(), " connected")
		go func() {
			defer func() {
				log.Infof("Closing connection %s", conn.RemoteAddr())
				conn.Close()
			}()
			r := bufio.NewReader(conn)
			w := bufio.NewWriter(conn)
			err := s.Handler(bufio.NewReadWriter(r, w))
			if err != nil {
				if err == io.EOF {
					return
				}
				log.Errorf("Handler returned with error: %v", err)
			}
		}()
	}
}
