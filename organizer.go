package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func organize() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path as an argument.")
		return
	}

	filePath := os.Args[1]
	//fileName := filepath.Base(filePath)
	currentFolder := filepath.Dir(filePath)

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	messages := string(content)

	svcID := extractValue(messages, `<ServiceID TYPE="UBYT">`, "</ServiceID>")
	svcName := extractValue(messages, `<ProtocolType TYPE="STR">`, "</ProtocolType>")

	messages = regexp.MustCompile(`^((?!MsgName).)*$`).ReplaceAllString(messages, "")
	messages = regexp.MustCompile(`<_MsgName TYPE="STR" NOXFER="TRUE">`).ReplaceAllString(messages, "")
	messages = regexp.MustCompile(`</_MsgName>`).ReplaceAllString(messages, "")
	messages = regexp.MustCompile(` `).ReplaceAllString(messages, "")
	messages = regexp.MustCompile(`\t`).ReplaceAllString(messages, "")
	messages = regexp.MustCompile(`\r\n\r\n`).ReplaceAllString(messages, "\r\n")
	for i := 0; i < 5; i++ {
		messages = regexp.MustCompile(`\n\n`).ReplaceAllString(messages, "\n")
	}
	messages = regexp.MustCompile(`^\n`).ReplaceAllString(messages, "")

	msgLines := strings.Split(messages, "\n")
	sort.Strings(msgLines)

	outputFileName := fmt.Sprintf("%s_%s.txt", svcID, svcName)
	output := outputFileName

	for i := 1; i < len(msgLines); i++ {
		output += fmt.Sprintf("\n%d: %s", i, msgLines[i])
	}

	err = ioutil.WriteFile(filepath.Join(currentFolder, outputFileName), []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
	}
}

func extractValue(content, startTag, endTag string) string {
	startPos := strings.Index(content, startTag) + len(startTag)
	endPos := strings.Index(content, endTag)
	return content[startPos:endPos]
}