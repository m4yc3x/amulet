# Amulet Project README

## Overview
The Amulet project is an emulation of the Wizard101 login server, written in Go. It aims to replicate the server's functionality to allow clients to connect, authenticate, and proceed to the game environment. This project is inspired by the Greyrose project, which was an initial attempt at emulating the Wizard101 login server but was acknowledged by its creator as a "quick and dirty" solution. Amulet seeks to provide a more structured and maintainable approach to server emulation.

## Credits
This project owes a great deal of its foundational knowledge to the research and documentation provided by the Greyrose project and the extensive findings of Joshsora, particularly in the realm of packet structure and message handling. For detailed packet information and DML (Data Markup Language) structure, Joshsora's repository (https://github.com/Joshsora/libki/wiki/) is an invaluable resource.

## Message Handling
Amulet utilizes the message information extracted from the game's `root.wad` file, as detailed in the Greyrose documentation. This information, once extracted, reveals the structure and content of messages that the server can send and receive. Each message is associated with a `ServiceID`, which determines the message file to use, and a `MessageID`, which is determined by sorting the message names alphabetically or by `MsgOrder` entries when available.

### Example
For handling login messages, Amulet references the `messages_login.go` file, which processes incoming messages based on their `ServiceID` and `MessageID`. The handling of these messages is crucial for authenticating users and transitioning them into the game.

```golang
func 7LoginMessages(data io.Reader, TCPPort int32, UDPPort int32, Key string) ([]byte, int, int) {}
```

## Packet Structure
Amulet adheres to the KingsIsle Networking Protocol (KINP) for message framing, which uses a hybrid delimited/length-prefixed approach. Each message begins with a "Start Signal" (`0xF00D`), followed by the length of the message payload and the payload itself. This structure is consistent across both TCP and UDP protocols.

## Contribution
Contributions to Amulet are welcome. Whether you're fixing bugs, adding new features, or improving documentation, your help is appreciated. Please refer to Joshsora's findings and the Greyrose project for guidance on the game's networking and message structures.

## License
Amulet is provided under an GPL license. Please see the LICENSE file for more details.

## Acknowledgments
- The Greyrose Project for the initial emulation attempt and documentation.
- Joshsora for the comprehensive analysis of Wizard101's networking and message structures.
- KingsIsle for greedily developing Wizard101, the game that inspired this project.