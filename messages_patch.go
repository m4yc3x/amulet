package main

import (
	"fmt"
	"io"
)

func _8PatchMessages(data io.Reader) []byte {
	msgid := readMSGID(data)
	//msglen := readUSHRT(data)

	switch msgid {
	case 1:
		fmt.Println("MSG: Latest File List")
		LatestVersion := readUINT(data)
		ListFileName := readSTR(data)
		ListFileType := readUINT(data)
		ListFileSize := readUINT(data)
		ListFileCRC := readUINT(data)
		ListFileURL := readSTR(data)
		URLPrefix := readSTR(data)
		URLSuffix := readSTR(data)
		fmt.Printf("Latest Version: %d\n", LatestVersion)
		fmt.Printf("File Name: %s\n", ListFileName)
		fmt.Printf("File Type: %d\n", ListFileType)
		fmt.Printf("File Size: %d\n", ListFileSize)
		fmt.Printf("File CRC: %d\n", ListFileCRC)
		fmt.Printf("File URL: %s\n", ListFileURL)
		fmt.Printf("URL Prefix: %s\n", URLPrefix)
		fmt.Printf("URL Suffix: %s\n", URLSuffix)
	case 2:
		fmt.Println("MSG: Latest File List V2")
		LatestVersion := readUINT(data)
		ListFileName := readSTR(data)
		ListFileType := readUINT(data)
		ListFileTime := readUINT(data)
		ListFileSize := readUINT(data)
		ListFileCRC := readUINT(data)
		ListFileURL := readSTR(data)
		URLPrefix := readSTR(data)
		URLSuffix := readSTR(data)
		Locale := readSTR(data)
		fmt.Printf("Latest Version: %d\n", LatestVersion)
		fmt.Printf("File Name: %s\n", ListFileName)
		fmt.Printf("File Type: %d\n", ListFileType)
		fmt.Printf("File Time: %d\n", ListFileTime)
		fmt.Printf("File Size: %d\n", ListFileSize)
		fmt.Printf("File CRC: %d\n", ListFileCRC)
		fmt.Printf("File URL: %s\n", ListFileURL)
		fmt.Printf("URL Prefix: %s\n", URLPrefix)
		fmt.Printf("URL Suffix: %s\n", URLSuffix)
		fmt.Printf("Locale: %s\n", Locale)
	case 3:
		fmt.Println("MSG: Next Version")
	}

	return nil
}