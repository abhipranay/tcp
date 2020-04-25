package protocol

import (
	"bufio"
	"encoding/binary"
	"errors"

	"abhi.com/tcp/message"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	// "google.golang.org/protobuf/proto"
)

func ReadRequest(r *bufio.Reader) (*message.RpcRequest, error) {
	req := &message.RpcRequest{}
	err := ReadPacket(req, r)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func ReadResponse(r *bufio.Reader) (*message.RpcResponse, error) {
	res := &message.RpcResponse{}
	err := ReadPacket(res, r)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ReadPacket(p proto.Message, r *bufio.Reader) error {
	var l [2]byte
	var i uint16 = 0
	for i < 2 {
		b, err := r.ReadByte()
		if err != nil {
			return err
		}
		l[i] = b
		i++
	}
	msgSize := binary.LittleEndian.Uint16([]byte{l[0], l[1]})
	log.Infof("Received message of length: %d", msgSize)
	i = 0
	pMsg := make([]byte, msgSize)
	for i < msgSize {
		n, err := r.Read(pMsg[i:])
		if err != nil {
			return err
		}
		i += uint16(n)
	}
	log.Infof("Raw message: %x", pMsg)
	proto.Unmarshal(pMsg, p)
	return nil
}

func WritePacket(m proto.Message, w *bufio.Writer) error {
	rMsg, err := proto.Marshal(m)
	if err != nil {
		log.Errorf("Failed to marshal packet for writing. %v", err)
		return err
	}
	l := binary.Size(rMsg)
	log.Infof("Len: %d", l)
	if l > 8190 {
		return errors.New("Message longer that 8190 bytes not allowed")
	}
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(l))
	log.Infof("Len bytes: % x", b)
	packet := make([]byte, l+2)
	copy(packet, b)
	copy(packet[2:], rMsg)
	log.Infof("Packet: % x", packet)
	tw := 0
	for tw != len(packet) {
		n, _ := w.Write(packet[tw:])
		w.Flush()
		tw += n
	}
	log.Infof("Sent message of length: %d", tw)
	return nil
}

func PrepareRequest(api string, m proto.Message) (*message.RpcRequest, error) {
	msg, err := proto.Marshal(m)
	if err != nil {
		log.Errorf("Failed to marshal payload. %v", err)
		return nil, err
	}
	rpc := &message.RpcRequest{
		Api:     &api,
		Payload: msg,
	}
	return rpc, nil
}

func PrepareResponse(m proto.Message) (*message.RpcResponse, error) {
	msg, err := proto.Marshal(m)
	if err != nil {
		return nil, err
	}
	r := &message.RpcResponse{
		Payload: msg,
	}
	return r, nil
}
