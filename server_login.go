package main

import (
	"fmt"
	"net"
	"time"
)

type LSSettings struct {
	TCPPort     int32
	UDPPort     int32
	Key         string
	IP          net.IP
	SocketOpen  bool
	SessionOpen bool
	SessionID   uint16
}

var lsSettings = LSSettings{
	TCPPort:     12000,
	UDPPort:     0,
	Key:         "m4yc3x",
	IP:          net.ParseIP("0.0.0.0"),
	SocketOpen:  false,
	SessionOpen: false,
	SessionID:   4513,
}

func LoginServer() {

	addr := fmt.Sprintf("%s:%d", lsSettings.IP, lsSettings.TCPPort)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("Login: Failed to start server: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("Login: Server started")

	for {
		fmt.Println("Login: Waiting for a connection...")
		lsSettings.SocketOpen = false
		lsSettings.SessionOpen = false

		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Login: Failed to accept connection: %v\n", err)
			continue
		}
		fmt.Println("Login: Connected!")

		go handleLoginConnection(conn)
	}
}

func handleLoginConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Login: AWAITING MESSAGE")

	buffer := make([]byte, 4096)

	for {
		if !lsSettings.SocketOpen {
			conn.SetReadDeadline(time.Now().Add(300 * time.Millisecond))
		} else {
			conn.SetReadDeadline(time.Now().Add(100000 * time.Millisecond))
		}

		n, err := conn.Read(buffer)
		fmt.Println(buffer)
		if err != nil {
			if err, ok := err.(net.Error); ok && err.Timeout() {
				fmt.Println("Login: NO DATA RECEIVED!")
				return
			}
			fmt.Printf("Login: Read error: %v\n", err)
			return
		}

		lsSettings.SocketOpen = true

		resp, r, l := Parse(buffer[:n], lsSettings.TCPPort, lsSettings.UDPPort, lsSettings.Key, lsSettings.SessionOpen, lsSettings.SessionID)
		if resp == nil {
			fmt.Printf("Login: Failed to parse message: %v\n", err)
			continue
		}

		response := resp
		requestLogin := r
		sessionOpen := l

		if response != nil {
			_, err := conn.Write(response)
			if err != nil {
				fmt.Printf("Login: Write error: %v\n", err)
				return
			}

			if requestLogin == 1 {
				packet := craftCharacterSelectedPacket()
				_, err := conn.Write(packet)
				if err != nil {
					fmt.Printf("Login: Write error: %v\n", err)
					return
				}
			}

			if sessionOpen == 1 {
				lsSettings.SessionOpen = true
			}
		}
	}
}

func craftCharacterSelectedPacket() []byte {
	packet := NewKIPacket()
	packet.Header(0x00, 0x00, 0x07, 0x03)
	packet._STR("127.0.0.1")
	packet._INT(gsSettings.TCPPort)
	packet._INT(gsSettings.UDPPort)
	packet._STR(gsSettings.Key)
	packet._GID(4295088136144)
	packet._GID(191965934135706025)
	packet._GID(123004564835992122)
	packet._STR("WizardCity/WC_Ravenwood")
	packet._STR("2572,4376,-28,5.55")
	packet._INT(0)
	packet._INT(0)
	packet._INT(0)
	packet._STR("Amulet.Login")
	return packet.Finalize()
}