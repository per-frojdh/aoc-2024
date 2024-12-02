package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	line, err := r.ReadString('\n')
	for err == nil {
		lines = append(lines, line)
		line, err = r.ReadString('\n')
	}

	if err != io.EOF {
		return false, []string{}
	}

	return true, lines
}
