package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"net"
)

func main() {
	for i:=0; i<10; i++ {
		go connServer()
	}
	select {
	}
}

func connServer() {
	conn, err  := net.Dial("tcp", ":9090")
	if err !=  nil {
		log.Error().Err(err).Msgf("connect to server failed")
	} else {
		log.Info().Msgf("connect to %s", conn.RemoteAddr())
		fmt.Fprint(conn, "hello, world")
		conn.Write([]byte("hello, world"))
		conn.Close()
	}
}
