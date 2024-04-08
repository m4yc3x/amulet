package main

import (
	"fmt"
	"io"
	"time"
)

func _1BaseMessages(data io.Reader, sessionOpen bool, sessionID uint16) ([]byte, bool) {
	msgid := readMSGID(data)
	//msglen := readUSHRT(data)

	if msgid == 1 {
		fmt.Println("MSG: Ping")

		if !sessionOpen {
			unixTimestamp := uint32(time.Now().Unix())

			packet := NewKIPacket()
			packet._UBYT(0x0d)
			packet._UBYT(0xf0)
			packet._UBYT(0x13)
			packet._UBYT(0x00)
			packet._UBYT(0x01)
			packet._UBYT(0x00)
			packet._UBYT(0x00)
			packet._UBYT(0x00)
			packet._USHRT(sessionID)
			packet._INT(0)
			packet._UINT(unixTimestamp)
			packet._INT(800)

			sessionOpen = true
			return packet.RawFinalize(), sessionOpen
		}

		resp := NewKIPacket()
		resp._UBYT(0x0D)
		resp._UBYT(0xF0)
		resp._USHRT(0x0900)
		resp._UBYT(0)
		resp._UBYT(0)
		resp._USHRT(0)
		resp._UBYT(1)
		resp._UBYT(2)
		resp._USHRT(4)

		return resp.RawFinalize(), sessionOpen
	} else if msgid == 2 {
		fmt.Println("MSG: Ping response")
	}

	return nil, sessionOpen
}