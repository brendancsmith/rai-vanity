package main

import (
	"os"
)

// AppendStringToFile :
//
// path: the path of the file
// text: the string to be appended. If you want to append text line to file,
//       add a newline '\n' to the string.
func AppendFile(path, text string) error {
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
