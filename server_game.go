package main

import (
	"fmt"
	"net"
)

type GS_Settings struct {
	TCPPort     int32
	UDPPort     int32
	IP          net.IP
	Key         string
	SocketOpen  bool
	SessionOpen bool
	SessionID   uint16
}

var gsSettings = GS_Settings{
	TCPPort: 		12170,
	UDPPort: 		12171,
	IP: 			net.ParseIP("0.0.0.0"),
	Key: 			"m4yc3x",
	SocketOpen: 	false,
	SessionOpen: 	false,
	SessionID: 		4513,
}

func GS() {
	addr := fmt.Sprintf("%s:%d", gsSettings.IP, gsSettings.TCPPort)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("GS: SocketException: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("GS: Waiting for a connection...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("GS: Accept error: %v\n", err)
			continue
		}

		go handleGameConnection(conn)
	}
}

func handleGameConnection(conn net.Conn) {
	defer conn.Close()

	bytes := make([]byte, 4096)
	for {
		bytesRead, err := conn.Read(bytes)
		if err != nil {
			fmt.Println("GS: NO DATA RECEIVED!")
			return
		}

		data := bytes[:bytesRead]
		resp, _, _ := Parse(data, gsSettings.TCPPort, gsSettings.UDPPort, gsSettings.Key, gsSettings.SessionOpen, gsSettings.SessionID)
		if resp != nil {
			conn.Write(resp)
		} else {
			hexData := fmt.Sprintf("%x", data)
			fmt.Printf("GS: Received: %s\n", hexData)
		}
	}
}