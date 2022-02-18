package main

import (
	"errors"
	"fmt"
	"runtime"
)

type Merror struct {
	msg    string
	inner  error
	caller string
}

func (m Merror) Error() string {
	return fmt.Sprintf("%s: %s %v", m.caller, m.msg, m.inner)
}

func New(msg string, inner error) Merror {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	var caller string
	if ok && details != nil {
		caller = details.Name()
	}
	return Merror{msg: msg, inner: inner, caller: caller}
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
		return New("opening file", err)
	}
	return nil
}

func openFile() error {

	return errors.New("File not found")
}
