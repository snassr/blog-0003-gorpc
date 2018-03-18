package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"reflect"

	"github.com/snassr/blog-0003-gorpc/egrpc"
)

func main() {
	// create person object
	person := &egrpc.Person{}

	// register person object as RPC (interface by the name `Person`)
	rpc.Register(person)

	// service address of server
	service := ":1200"

	// create tcp address
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		log.Fatal(err)
	}

	// tcp network listener
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// handle tcp client connections
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listener accept error:", err)
		}

		// print connection info
		fmt.Println("received message", reflect.TypeOf(conn), conn)

		// handle client connections via rpc
		rpc.ServeConn(conn)
	}
}
