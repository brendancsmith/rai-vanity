package main

import (
	"fmt"
	"os"
)

func append(path, text string) error {
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

func appendLn(path, text, sep string) error {
	return append(path, fmt.Sprintf("%s%s", text, sep))
}
