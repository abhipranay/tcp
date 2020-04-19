package client

import (
	"bufio"
	"net"

	"abhi.com/tcp/protocol"
	log "github.com/sirupsen/logrus"

	// "google.golang.org/protobuf/proto"
	"github.com/golang/protobuf/proto"
)

type Client struct {
	Proto string
	SAddr string
	conn  *net.TCPConn
	rw    *bufio.ReadWriter
}

func (c *Client) Connect() error {
	srvTcpAddr, err := net.ResolveTCPAddr(c.Proto, c.SAddr)
	if err != nil {
		log.Panicf("Failed to resolve server address %s. %v", c.SAddr, err)
	}
	conn, err := net.DialTCP(c.Proto, nil, srvTcpAddr)
	if err != nil {
		log.Panicf("Failed to connect to server at address %s. %v", srvTcpAddr.String(), err)
		return err
	}
	c.conn = conn
	c.rw = bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	return nil
}

func (c *Client) Send(m proto.Message) error {
	req, _ := protocol.PrepareRequest("login", m)
	err := protocol.WritePacket(req, c.rw.Writer)
	if err != nil {
		log.Errorf("Client failed to send message: %v", err)
		return err
	}
	res, err := protocol.ReadResponse(c.rw.Reader)
	if err != nil {
		log.Errorf("Failed to read server response: %v", err)
		return err
	}
	log.Infof("Server: %x", res)
	return nil
}
