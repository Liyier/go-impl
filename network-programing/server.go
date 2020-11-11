package main

import (
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	server, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Error().Err(err).Msg("server bind address failed")
		os.Exit(1)
	} else {
		log.Info().Msgf("listen addr : %s", server.Addr())
	}
	defer server.Close()
	var connChan = make(chan net.Conn, 10)
	for i:=0; i<10; i++ {
		go func() {
			for {
				select {
				case conn := <- connChan:
					handleConn(conn)
				}
			}
		}()
	}
	for {
		if conn, err := server.Accept(); err == nil {
			connChan <- conn
		}
	}
}

func handleConn(conn net.Conn)  {
	defer conn.Close()
	if buffer, err := ioutil.ReadAll(conn); err != nil {
		log.Error().Err(err).Msgf("read from conn err, remote addr:%s", conn.RemoteAddr())
	} else {
		log.Info().Bytes("msg", buffer).Msgf("receve msg from %s", conn.RemoteAddr())
		conn.Write(buffer)
	}
}
