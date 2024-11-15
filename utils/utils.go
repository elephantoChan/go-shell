package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Echo(Words ...string) string {
	var ToPrint string
	for _, Word := range Words {
		ToPrint += Word + " "
	}
	return strings.TrimSpace(ToPrint)
}

func Cat(FileName string) (string, error) {
	file, err := os.ReadFile(FileName)
	if err != nil {
		fmt.Print("error: filename is incorrect; can't find ", FileName)
		return "", err
	}
	return string(file), nil
}

func Grep(Content []byte, regex string) []string {
	re, err := regexp.Compile(regex)
	if err != nil {
		fmt.Print(err)
	}
	var b []string
	for _, line := range strings.Split(string(Content), "\n") {
		match := re.FindAll([]byte(line), -1)
		if match != nil {
			b = append(b, line)
		}
	}
	return b
}
