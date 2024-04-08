# Amulet Project README

## Overview
The Amulet project is an emulation of the Wizard101 login server, written in Go. It aims to replicate the server's functionality to allow clients to connect, authenticate, and proceed to the game environment. This project is inspired by the Greyrose project, which was an initial attempt at emulating the Wizard101 login server but was acknowledged by its creator as a "quick and dirty" solution. Amulet seeks to provide a more structured and maintainable approach to server emulation.

## Credits
This project owes a great deal of its foundational knowledge to the research and documentation provided by the Greyrose project and the extensive findings of Joshsora, particularly in the realm of packet structure and message handling. For detailed packet information and DML (Data Markup Language) structure, Joshsora's repository (https://github.com/Joshsora/libki/wiki/) is an invaluable resource.

## Message Handling
KingsIsle, the developers behind Wizard101, have not prioritized security in certain aspects of their game development. This oversight has left a wealth of message information in clear text within the game files, which can be utilized for our purposes.

## Extracting Messages

To access the current messages, one needs to extract them from the `root.wad` file found in the path `Wizard101\Data\GameData\`.

### Required Tools

- **Quickbms**: A free tool available at [Quickbms official site](https://aluigi.altervista.org/quickbms.htm).
- **Kiwad Extractor**: A specific script for Quickbms that can be downloaded from [Kiwad Extractor](https://aluigi.altervista.org/bms/wizard101_kiwad.bms).

### Extraction Process

1. Download both Quickbms and the Kiwad Extractor script.
2. Extract all files from the Quickbms zip archive.
3. Place the `wizard101_kiwad.bms` file in the same directory as the extracted Quickbms files.
4. Drag and drop `wizard101_kiwad.bms` onto `quickbms.exe`.
5. In the prompted window, select your `Root.wad` file and choose a destination folder for the extracted contents.

After Quickbms completes the extraction process, navigate to the chosen directory. You will find numerous `.xml` files, with those ending in `Messages.xml` being of particular interest. Additionally, some message files can be found within the `Messages` sub-folder.

## Understanding Message Files

Each message file contains a `ServiceID` entry, which is crucial for determining the appropriate message file to use based on the packet's Service ID. For example, a packet with a Service ID of 5 would correspond to the `GameMessages.xml` file, as it has a Service ID of 5.

### MessageID

The determination of `MessageID` is slightly more complex. In the absence of `MsgOrder` entries (which is common), the `MessageID` is assigned by sorting the message names alphabetically.

### Tool for Message Sorting


## Packet Structure
Amulet adheres to the KingsIsle Networking Protocol (KINP) for message framing, which uses a hybrid delimited/length-prefixed approach. Each message begins with a "Start Signal" (`0xF00D`), followed by the length of the message payload and the payload itself. This structure is consistent across both TCP and UDP protocols. For comprehensive packet information, the repository maintained by Joshsora is highly recommended: [Joshsora's libki wiki](https://github.com/Joshsora/libki/wiki/). The findings within have been instrumental in understanding the intricacies of packet handling in KingsIsle games, saving countless hours of research. For redundancy and ease of access, the key information has been summarized below.

## DML (Data Markup Language)

DML is a data serialization system utilized in KingsIsle games, such as Wizard101 and Pirate101. It defines application-specific data structures that are often sent over a session of the KingsIsle Networking Protocol (KINP). It's important to note that DML is separate from KINP, serving as the application data for these games.

### DML Record

A DML Record comprises multiple DML Fields, each containing a field name (identifier), data type, value, and a "transfer" property indicating if the field should be serialized into the binary representation.

### Usage

DML is primarily used within XML Message Modules, defining DML structures for messages sent over a KINP session. It's also found in configuration files, such as `defaultconfig.xml` within the `Root.wad` of Wizard101.

## DML Syntax

DML Records are represented as XML in Message Modules and various configuration files. A Record starts with a `<RECORD>` element, with each child element defining a Field. Fields must define a data type and may optionally specify a NOXFER attribute to control serialization.

## DML Data Types

DML supports a variety of data types, from basic integers and strings to floating-point numbers and globally unique identifiers (GIDs), all with specific binary representations.

## KingsIsle Networking Protocol (KINP)

KINP employs a hybrid delimited/length-prefixed approach for message framing. Each message starts with a "Start Signal" (`0xF00D`), followed by the message length and payload. This structure is consistent across TCP and UDP protocols.

### Message Structure

Messages include a start signal, length, and payload, with the payload structure varying based on whether it's a control message or application-specific data.

## Opcodes

Opcodes define the type of message being sent or received, such as Session Offer, Session Accept, and Keep Alive messages. Each has a specific structure and purpose within the session lifecycle.

## Message Modules

Message Modules are XML-based resources defining a Message Service, containing multiple DML Records as templates for messages. These can be referenced by a unique ServiceID, with each record given a MessageType based on a _MsgOrder field or alphabetical ordering.

### XML Format

The root element of an XML message module is often the file name without the `.xml` extension. Special reserved field names include _MsgName, _MsgDescription, _MsgHandler, _MsgAccessLvl, and _MsgOrder, providing metadata for messages.

### XML Example

```xml
<RECORD>
  <_Opcode TYPE="USHRT" NOXFER="TRUE">1</_Opcode>
  <Name TYPE="STR"></Name>
  <Age TYPE="UBYT"></Age>
</RECORD>
```

## DML Message Structure

When a DML message is transferred within a KINP message, a header indicates the Service ID and Message Type for deserializing the message into a Record. The binary serialization follows the message template structure.


# Contribution
Contributions to Amulet are welcome. Whether you're fixing bugs, adding new features, or improving documentation, your help is appreciated. Please refer to Joshsora's findings and the Greyrose project for guidance on the game's networking and message structures.

# License
Amulet is provided under an GPL license. Please see the LICENSE file for more details.

# Acknowledgments
- The Greyrose Project for the initial emulation attempt and documentation.
- Joshsora for the comprehensive analysis of Wizard101's networking and message structures.
- KingsIsle for greedily developing Wizard101, the game that inspired this project.