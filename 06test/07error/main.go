package main

import (
	"errors"
	"os"
)

func main() {
}

func LestLookError(filename string) error {
	_, err := os.ReadFile(filename)
	return err
}

func NewMyError()error{
	return errors.New("Hello")
}