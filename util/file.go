package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadInputIntoLines(filename string) (bool, []string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return false, []string{}
	}
	defer file.Close()

	r := bufio.NewReader(file)
	var lines []string

	for {
		line, err := r.ReadString('\n')
		line = strings.TrimRight(line, "\r\n") // Remove trailing newline characters
		lines = append(lines, line)

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return false, []string{}
		}
	}

	// Remove any empty line caused by appending at EOF
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return true, lines
}

func ReadInputIntoString(filename string) (bool, string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	defer file.Close()

	r := bufio.NewReader(file)
	var builder strings.Builder

	for {
		line, err := r.ReadString('\n')
		builder.WriteString(strings.TrimRight(line, "\r\n"))

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return false, ""
		}
	}

	return true, builder.String()
}
