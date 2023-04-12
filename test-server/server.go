package main

import (
	"net"
	"net/rpc"
	"os"
	"strings"

	goridgeRpc "github.com/spiral/goridge/v3/pkg/rpc"
)

// Service
type Service struct{}

// Payload
type Payload struct {
	Name  string
	Value int
	Keys  map[string]string
}

// Negate number
func (s *Service) Negate(i int64, r *int64) error {
	*r = -i
	return nil
}

// Ping pong
func (s *Service) Ping(msg string, r *string) error {
	if msg == "ping" {
		*r = "pong"
	}
	return nil
}

// Echo returns incoming message
func (s *Service) Echo(msg string, r *string) error {
	*r = msg
	return nil
}

// Process performs payload conversion
func (s *Service) Process(msg Payload, r *Payload) error {
	r.Name = strings.ToUpper(msg.Name)
	r.Value = -msg.Value

	if len(msg.Keys) != 0 {
		r.Keys = make(map[string]string)
		for n, v := range msg.Keys {
			r.Keys[v] = n
		}
	}

	return nil
}

// EchoBinary work over binary data
func (s *Service) EchoBinary(msg []byte, out *[]byte) error {
	*out = append(*out, msg...)

	return nil
}

func main() {
	var ln net.Listener
	var err error
	if len(os.Args) == 2 {
		ln, err = net.Listen("unix", os.Args[1])
	} else {
		ln, err = net.Listen("tcp", ":7079")
	}

	if err != nil {
		panic(err)
	}

	err = rpc.Register(new(Service))
	if err != nil {
		panic(err)
	}

	for {
		conn, err2 := ln.Accept()
		if err2 != nil {
			continue
		}
		go rpc.ServeCodec(goridgeRpc.NewCodec(conn))
	}
}
