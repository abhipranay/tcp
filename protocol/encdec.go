package protocol

import (
	"bufio"
	"encoding/binary"
	"errors"

	"abhi.com/tcp/message"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func ReadRequest(r *bufio.Reader) (proto.Message, error) {
	var l [2]byte
	var i uint16 = 0
	for i < 2 { 
		b, err := r.ReadByte()
		if err != nil {
			log.Errorf("Error while reading: %v", err)
			return nil, err
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
			log.Errorf("Error while reading: %v", err)
			return nil, err
		}
		i += uint16(n)
	}
	log.Infof("Raw message: %x", pMsg)
	req := &message.RpcRequest{}
	proto.Unmarshal(pMsg, req)
	return req, nil
}

func WriteResponse(m proto.Message, w *bufio.Writer) error {
	msg, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	rMsg, err := proto.Marshal(&message.Response{
		Payload: msg,
	})
	if err != nil {
		log.Errorf("Failed to send resposne: %v", err)
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
	packet := make([]byte, l + 2)
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

func PrepareRequest(m proto.Message) (proto.Message, error) {
	msg, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	rMsg, err := proto.Marshal(&message.RpcRequest{
		Payload: msg,
	})
}