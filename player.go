package main

type PlayerStruct struct {
	X         float32
	Y         float32
	Z         float32
	Rot       float32
	GID       int64
	ZoneName  string
	ZoneID    int64
	MarkerX   uint16
	MarkerY   uint16
	MarkerZ   uint16
	MarkerRot uint8
}

var player1 PlayerStruct
