package main

import (
	"errors"
	"fmt"
)

type Merror struct {
	msg   string
	inner error
}

func (m Merror) Error() string {
	return fmt.Sprintf("%s: %v", m.msg, m.inner)
}

func New(msg string, inner error) Merror {
	return Merror{msg: msg, inner: inner}
}

func main() {
	err := processFile()
	if err != nil {
		fmt.Println(err)
	}

}

func processFile() error {

	err := openFile()
	if err != nil {
		return New("processFile opening file", err)
	}
	return nil
}

func openFile() error {

	return errors.New("File not found")
}