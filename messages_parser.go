package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func Parse(bytess []byte, TCPPort int32, UDPPort int32, Key string, SessionOpen bool, SessionID uint16) ([]byte, int, int) {
	ByteStream := bytes.NewReader(bytess)
	var header [2]byte
	ByteStream.Read(header[:])

	if fmt.Sprintf("%02X", header) != "0DF0" {
		fmt.Println("UNKNOWN PACKET TYPE!")
		return nil, 0, 0
	}

	var msglen uint16
	binary.Read(ByteStream, binary.LittleEndian, &msglen)
	opcode := readUBYT(ByteStream)
	svcid := readUBYT(ByteStream)
	msgid := readUBYT(ByteStream)

	fmt.Printf("SVCID: %d, MSGID: %d\n", svcid, msgid)

	switch svcid {
	case 0:
		return _0ControlMessages(ByteStream, int(opcode), int(msglen), bytess)
	case 1:
		resp, r := _1BaseMessages(ByteStream, SessionOpen, SessionID)
		r = r
		return resp, 0, 0
	case 5:
		return _5GameMessages(ByteStream, TCPPort, UDPPort, Key)
	case 7:
		return _7LoginMessages(ByteStream, TCPPort, UDPPort, Key)
	case 8:
		resp := _8PatchMessages(ByteStream)
		return resp, 0, 0
	case 52:
		fmt.Println("Quest Message. No current handlers")
	default:
		fmt.Println("UNHANDLED MESSAGE!")
		hexbytes := ""
		for i := 0; i < int(msglen)+4; i++ {
			hexbytes += fmt.Sprintf("%02X", bytess[i])
		}
		fmt.Println(hexbytes)
	}

	return nil, 0, 0
}