package server

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"

	pb "github.com/murtaza-u/rtspproxy/proto/gen/go/proxy"

	"github.com/aler9/gortsplib/v2"
	"github.com/aler9/gortsplib/v2/pkg/media"
	"github.com/pion/rtp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

type Packet struct {
	Medi *media.Media
	Pkt  *rtp.Packet
}

type ProxySrv struct {
	stream *gortsplib.ServerStream

	pb.UnimplementedProxySvcServer
}

func (s *ProxySrv) Stream(stream pb.ProxySvc_StreamServer) error {
	for {
		encP, err := stream.Recv()
		if err != nil {
			return err
		}
		p, err := decodePacket(encP.GetData())
		if err != nil {
			return grpc.Errorf(codes.InvalidArgument,
				"failed to decode packet: %s", err.Error())
		}

		s.stream.WritePacketRTP(p.Medi, p.Pkt)
	}
}

func decodePacket(data []byte) (*Packet, error) {
	p := new(Packet)
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

const addr = ":5000"

func Start() error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	s := grpc.NewServer()
	pb.RegisterProxySvcServer(s, &ProxySrv{})
	reflection.Register(s)

	return s.Serve(ln)
}
