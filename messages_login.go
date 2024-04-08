package main

import (
	"fmt"
	"io"
)

func _7LoginMessages(data io.Reader, TCPPort int32, UDPPort int32, Key string) ([]byte, int, int) {
	msgid := readMSGID(data)
	//msglen := USHRT(data)

	switch msgid {
	case 1:
		fmt.Println("MSG_CHARACTERINFO")
	case 2:
		fmt.Println("MSG_CHARACTERLIST")
	case 3:
		fmt.Println("MSG_CHARACTERSELECTED")
	case 4:
		fmt.Println("MSG_CREATECHARACTER")
	case 5:
		fmt.Println("MSG_DELETECHARACTER")
	case 6:
		fmt.Println("MSG_CHARACTERCREATED")
	case 7:
		fmt.Println("MSG_CHARACTERDELETED")
	case 8:
		fmt.Println("MSG_CHARACTERCREATIONFAILED")
	case 9:
		fmt.Println("MSG_CHARACTERDELETIONFAILED")
	case 10:
		fmt.Println("MSG_CHARACTERNAMEINVALID")
	case 11:
		fmt.Println("MSG_CHARACTERNAMETAKEN")
	case 12:
		fmt.Println("MSG_STARTCHARACTERLIST")
	case 13:
		fmt.Println("MSG_ENDCHARACTERLIST")
	case 14:
		fmt.Println("MSG_CHARACTERLISTFULL")
	case 15:
		fmt.Println("MSG_CHARACTERLISTEMPTY")
	case 16:
		fmt.Println("MSG_CHARACTERLISTINVALID")
	case 17:
		fmt.Println("MSG_CHARACTERLISTUPDATED")
	case 18:
		fmt.Println("MSG_CHARACTERLISTUPDATEFAILED")
	case 19:
		fmt.Println("MSG_CHARACTERLISTUPDATESUCCESSFUL")
	case 20:
		fmt.Println("MSG_CHARACTERLISTUPDATECANCELLED")
	case 21:
		fmt.Println("MSG_USER_AUTHEN")
	case 22:
		fmt.Println("MSG_USER_AUTHEN_V2")
	case 23:
		fmt.Println("MSG_SAVECHARACTER")
	case 24:
		fmt.Println("MSG_WEB_AUTHEN")
	case 25:
		fmt.Println("MSG_WEB_VALIDATE")
	case 26:
		fmt.Println("MSG_CHANGECHARACTERNAME")
	case 27:
		fmt.Println("MSG_USER_AUTHEN_V3")
	default:
		fmt.Println("UNSUPPORTED MESSAGE! Make sure you are running revision r667549.Wizard_1_390")
	}

	return nil, 0, 0
}