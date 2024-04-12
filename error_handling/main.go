package main

import (
	"errors"
	"fmt"
)

// to create an error simply make a call to error.New(string).
func main() {
	err := doesReturnError1()
	if err != nil {
		fmt.Printf("err1: %v\n", err)
	}
	err2 := doesReturnError2(errors.New("Error created"))
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2.Error())
	}
}

func doesReturnError1() error {
	err := errors.New("Error created")
	return err
}
func doesReturnError2(err error) error {
	return err
}
