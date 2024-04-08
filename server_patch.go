package main

import (
	"fmt"
	"net"
)

type PS_Settings struct {
	TCPPort     int32
	UDPPort     int32
	Key         string
	IP          net.IP
	SocketOpen  bool
	SessionOpen bool
	SessionID   uint16
}

var psSettings = PS_Settings{
	TCPPort: 12500,
	UDPPort: 0,
	Key: "m4yc3x",
	IP: net.ParseIP("0.0.0.0"),
	SocketOpen: false,
	SessionOpen: false,
	SessionID: 4513,
}

func PS() {
	addr := fmt.Sprintf("%s:%d", psSettings.IP, psSettings.TCPPort)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("Patch: SocketException: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("Patch: Waiting for a connection...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Patch: Accept error: %v\n", err)
			continue
		}

		go handlePatchConnection(conn)
	}
}

func handlePatchConnection(conn net.Conn) {
	defer conn.Close()

	bytes := make([]byte, 4096)
	for {
		bytesRead, err := conn.Read(bytes)
		if err != nil {
			fmt.Println("Patch: NO DATA RECEIVED!")
			return
		}

		data := bytes[:bytesRead]
		resp, _, _ := Parse(data, psSettings.TCPPort, psSettings.UDPPort, psSettings.Key, psSettings.SessionOpen, psSettings.SessionID)
		if resp != nil {
			conn.Write(resp)
		} else {
			hexData := fmt.Sprintf("%x", data)
			fmt.Printf("Patch: Received: %s\n", hexData)
		}
	}
}