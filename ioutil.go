package main

import (
	"bufio"
	"io"
	"os"
)

func readLines(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return lines
}

// append :
//
// path: the path of the file
// text: the string to be appended. If you want to append text line to file,
//       add a newline '\n' to the string.
func fileAppend(path, text string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}
