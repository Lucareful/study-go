package main

import (
	"fmt"
	"github.com/pkg/errors"
)

var myErr = errors.New("something not works.")

func main() {
	err := test3()
	fmt.Printf("main: %+v\n", err)
}

func test1() error {
	return errors.Wrapf(myErr, "test1 failed.")
}

func test2() error {
	return test1()
}

func test3() error {
	return test2()
}
