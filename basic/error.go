package main

import (
	"log"
	"os"
)

// Assume we have a custom error type defined
type MyError struct {
	msg string
}

// Error implements the error interface for MyError
func (e *MyError) Error() string {
	return e.msg
}

func otherFunc() (int, error) {
	// This function is assumed to return an error or nil
	return 0, &MyError{msg: "Custom error occurred"}
}

func errors() {
	// Attempt to open a file
	f, err := os.Open("/home/somaz/PrivateWork/go-study/basic")
	if err != nil {
		// If there is an error, log it and exit
		log.Fatal(err.Error())
	}
	// If no error, print the file name
	println(f.Name())

	// Now, let's use the otherFunc and handle different error types
	_, err = otherFunc()
	switch err.(type) {
	default: // no error
		println("ok")
	case *MyError: // custom error type
		log.Print("Log my error: ", err.Error())
	case error: // general error case
		log.Fatal(err.Error())
	}
}
