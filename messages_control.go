package main

import (
	"fmt"
	"io"
)

func _0ControlMessages(data io.Reader, opCode int, length int, bytes []byte) ([]byte, int, int) {
	if opCode == 0 {
		fmt.Println("OPCODE: SESSION OFFER")
		sessionid := readUSHRT(data)
		undefined := readUINT(data)
		timestamp := readINT(data)
		milliseconds := readUINT(data)
		fmt.Printf("Session ID: %d\n", sessionid)
		fmt.Printf("undefined: %d\n", undefined)
		fmt.Printf("Timestamp: %d\n", timestamp)
		fmt.Printf("Milliseconds: %d\n", milliseconds)
	} else if opCode == 3 {
		fmt.Println("OPCODE: KEEP ALIVE")
		sessionid := readUSHRT(data)
		milliseconds := readUSHRT(data)
		minutes := readUSHRT(data)
		fmt.Printf("Sessiond ID: %d\n", sessionid)
		fmt.Printf("Milliseconds: %d\n", milliseconds)
		fmt.Printf("Minutes: %d\n", minutes)

		b := make([]byte, length+4)
		copy(b, bytes[:length+4])

		b[5] = 0x04
		return b, 0, 0
	} else if opCode == 4 {
		fmt.Println("OPCODE: KEEP ALIVE RESPONSE")
		undefined := readUSHRT(data)
		timestamp := readUINT(data)
		fmt.Printf("Unknown: %d\n", undefined)
		fmt.Printf("Timestamp: %d\n", timestamp)
	} else if opCode == 5 {
		fmt.Println("OPCODE: SESSION ACCEPT")
		undefined := readUSHRT(data)
		undefined2 := readUINT(data)
		timestamp := readUINT(data)
		milliseconds := readUINT(data)
		sessionID := readUSHRT(data)
		fmt.Printf("Unknown1: %d\n", undefined)
		fmt.Printf("Unknown2: %d\n", undefined2)
		fmt.Printf("Timestamp: %d\n", timestamp)
		fmt.Printf("Milliseconds: %d\n", milliseconds)
		fmt.Printf("Session ID: %d\n", sessionID)
		return nil, 0, 1
	} else {
		fmt.Println("UNSUPPORTED OPCODE!")
	}

	return nil, 0, 0
}