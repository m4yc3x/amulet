package main

import (
	"fmt"
	"io"
	"math"
)

func _5GameMessages(data io.Reader, TCPPort int32, UDPPort int32, Key string) ([]byte, int, int) {
	msgid := readMSGID(data)
	//msglen := readUSHRT(data)

	if msgid == 2 {
		fmt.Println("MSG_CHARACTERLIST")
		return nil, 0, 0
	} else if msgid == 36 {
		fmt.Println("MSG_CLIENTMOVE")
		x := float32(readSHRT(data) * 4)
		y := float32(readSHRT(data) * 4)
		z := float32(readSHRT(data) * 4)
		rot := float32(readUBYT(data)) * float32(math.Pi) * 2 / 250
		fmt.Printf("X: %f\nY: %f\nZ: %f\nRot: %f\n", x, y, z, rot)
		fmt.Println("Zone:", readUBYT(data))

		// PlayerData updates here

		// Uncomment to add marker trails

		// Save player location to marker variables

		return nil, 0, 0

	} else if msgid == 37 {
		fmt.Println("MSG_RECMSG_CLIENTMOVESTATEALL")

		return nil, 0, 0

	} else if msgid == 40 {
		fmt.Println("MSG_CLIENT_DISCONNECT")

		return nil, 0, 0
	} else if msgid == 100 {
		fmt.Println("MGS_JUMP")
	
		return nil, 0, 0
	} else if msgid == 110 {
		fmt.Println("MSG_MARK_LOCATION")
	
		return nil, 0, 0
	} else if msgid == 171 {
		fmt.Println("MSG_RECALL")

		// KIPacket creation and data handling

		return nil, 0, 0
	} else if msgid <= 253 {
		fmt.Println("Unhandled message!")
		fmt.Printf("SERVICE: 5 (GameMessages)\nMESSAGE ID: %d\n", msgid)
		return nil, 0, 0
	}

	fmt.Println("UNKNOWN MESSAGE ID! Make sure you are running revision r667549.Wizard_1_390")
	return nil, 0, 0
}