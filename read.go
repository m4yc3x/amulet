package main

import (
	"encoding/binary"
	"io"
)

func readMSGID(data io.Reader) uint8 {
	var msgid uint8
	binary.Read(data, binary.LittleEndian, &msgid)
	return msgid
}

func readBYT(data io.Reader) int8 {
	var val int8
	binary.Read(data, binary.LittleEndian, &val)
	return val
}

func readUBYT(data io.Reader) uint8 {
	var val uint8
	binary.Read(data, binary.LittleEndian, &val)
	return val
}

func readSHRT(data io.Reader) int16 {
	var val int16
	binary.Read(data, binary.LittleEndian, &val)
	return val
}

func readUSHRT(data io.Reader) uint16 {
	var val uint16
	binary.Read(data, binary.LittleEndian, &val)
	return val
}

func readINT(data io.Reader) int32 {
	var val int32
	binary.Read(data, binary.LittleEndian, &val)
	return val
}

func readUINT(data io.Reader) uint32 {
	var val uint32
	binary.Read(data, binary.LittleEndian, &val)
	return val
}

func readSTR(data io.Reader) string {
	length := readUSHRT(data)
	buf := make([]byte, length)
	data.Read(buf)
	return string(buf)
}

func readWSTR(data io.Reader) string {
	return readSTR(data)
}

func readFLT(data io.Reader) float32 {
	var val float32
	binary.Read(data, binary.LittleEndian, &val)
	return val
}

func readDBL(data io.Reader) float64 {
	var val float64
	binary.Read(data, binary.LittleEndian, &val)
	return val
}

func readGID(data io.Reader) uint64 {
	var val uint64
	binary.Read(data, binary.LittleEndian, &val)
	return val
}