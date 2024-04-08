package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"strings"
)

func writeBYT(message *bytes.Buffer, data int8) {
	binary.Write(message, binary.LittleEndian, data)
}

func writeUBYT(message *bytes.Buffer, data uint8) {
	binary.Write(message, binary.LittleEndian, data)
}

func writeSHRT(message *bytes.Buffer, data int16) {
	binary.Write(message, binary.LittleEndian, data)
}

func writeUSHRT(message *bytes.Buffer, data uint16) {
	binary.Write(message, binary.LittleEndian, data)
}

func writeINT(message *bytes.Buffer, data int32) {
	binary.Write(message, binary.LittleEndian, data)
}

func writeUINT(message *bytes.Buffer, data uint32) {
	binary.Write(message, binary.LittleEndian, data)
}

func writeSTR(message *bytes.Buffer, data string) {
	length := uint16(len(data))
	binary.Write(message, binary.LittleEndian, length)
	message.WriteString(data)
}

func writeWSTR(message *bytes.Buffer, data string) {
	writeSTR(message, data)
}

func writeFLT(message *bytes.Buffer, data float32) {
	binary.Write(message, binary.LittleEndian, data)
}

func writeDBL(message *bytes.Buffer, data float64) {
	binary.Write(message, binary.LittleEndian, data)
}

func writeGID(message *bytes.Buffer, data uint64) {
	binary.Write(message, binary.LittleEndian, data)
}

type KIPacket struct {
	buffer   *bytes.Buffer
	outerLen int
	innerLen uint16
}

func NewKIPacket() *KIPacket {
	return &KIPacket{
		buffer: new(bytes.Buffer),
	}
}

func (p *KIPacket) Header(control, opcode, svcid, msgid uint8) {
	writeUBYT(p.buffer, 0x0D)
	writeUBYT(p.buffer, 0xF0)
	writeUSHRT(p.buffer, 0)
	writeUBYT(p.buffer, control)
	writeUBYT(p.buffer, opcode)
	writeUSHRT(p.buffer, 0)
	writeUBYT(p.buffer, svcid)
	writeUBYT(p.buffer, msgid)
	writeUSHRT(p.buffer, 0)
	p.outerLen -= 5
	p.innerLen -= 10
}

func (p *KIPacket) _BYT(data int8) {
	writeBYT(p.buffer, data)
	p.innerLen++
	p.outerLen++
}

func (p *KIPacket) _UBYT(data uint8) {
	writeUBYT(p.buffer, data)
	p.innerLen++
	p.outerLen++
}

func (p *KIPacket) _SHRT(data int16) {
	writeSHRT(p.buffer, data)
	p.innerLen += 2
	p.outerLen += 2
}

func (p *KIPacket) _USHRT(data uint16) {
	writeUSHRT(p.buffer, data)
	p.innerLen += 2
	p.outerLen += 2
}

func (p *KIPacket) _INT(data int32) {
	writeINT(p.buffer, data)
	p.innerLen += 4
	p.outerLen += 4
}

func (p *KIPacket) _UINT(data uint32) {
	writeUINT(p.buffer, data)
	p.innerLen += 4
	p.outerLen += 4
}

func (p *KIPacket) _STR(data string) {
	writeSTR(p.buffer, data)
	p.innerLen += uint16(len(data) + 2)
	p.outerLen += len(data) + 2
}

func (p *KIPacket) _WSTR(data string) {
	writeWSTR(p.buffer, data)
	p.innerLen += uint16(len(data) + 2)
	p.outerLen += len(data) + 2
}

func (p *KIPacket) _FLT(data float32) {
	writeFLT(p.buffer, data)
	p.innerLen += 4
	p.outerLen += 4
}

func (p *KIPacket) _DBL(data float64) {
	writeDBL(p.buffer, data)
	p.innerLen += 8
	p.outerLen += 8
}

func (p *KIPacket) _GID(data uint64) {
	writeGID(p.buffer, data)
	p.innerLen += 8
	p.outerLen += 8
}

func (p *KIPacket) _HEXSTRING(data string) {
	tempStr := strings.Join(strings.Fields(data), "")
	temp, _ := hex.DecodeString(tempStr)
	p.buffer.Write(temp)
	p.innerLen += uint16(len(temp))
	p.outerLen += len(temp)
}

func (p *KIPacket) Finalize() []byte {
	p._BYT(0)
	p.innerLen++
	p.outerLen++
	output := make([]byte, p.buffer.Len())
	copy(output, p.buffer.Bytes())
	binary.LittleEndian.PutUint16(output[2:], uint16(p.outerLen))
	binary.LittleEndian.PutUint16(output[10:], p.innerLen)
	return output
}

func (p *KIPacket) RawFinalize() []byte {
	p._BYT(0)
	p.innerLen++
	p.outerLen++
	return p.buffer.Bytes()
}