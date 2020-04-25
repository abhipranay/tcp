package protocol

import (
	"bufio"

	"abhi.com/tcp/message"
	log "github.com/sirupsen/logrus"

	// "google.golang.org/protobuf/proto"
	"github.com/golang/protobuf/proto"
)

func ProtobufMessageHandler(rw *bufio.ReadWriter) error {
	r := rw.Reader
	w := rw.Writer

	for {
		pMsg, err := ReadRequest(r)
		if err != nil {
			return err
		}
		log.Infof("Proto Message: %x", pMsg)
		res := getLoginResponse()
		err = WritePacket(res, w)
		if err != nil {
			return err
		}
	}
}

func getLoginResponse() proto.Message {
	t := "abcdefghijk"
	r := "reftok"
	return &message.LoginResponse{
		Token:        &t,
		RefreshToken: &r,
	}
}
