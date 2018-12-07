package main

import (
	"errors"
	"fmt"
	"os"
)

/*Anthonys-MacBook-Pro:go mcclayac$ godoc fmt Errorf
func Errorf(format string, a ...interface{}) error
Errorf formats according to a format specifier and returns the string as
a value that satisfies error.
*/

var (
	errorEmptyString = errors.New("Unwilling to print an empty string #2")
)  // error package

func myprinter(msg string) error {
	if msg == "" {
		//return fmt.Errorf("Unwilling to print an empty string")
		//return errorEmptyString
		panic(errorEmptyString)   // not for general usage
	}

	_, err := fmt.Printf("%s\n", msg)
	return err
}

func main() {
	if err := myprinter("The Nuts!"); err != nil {
		fmt.Printf("Printer FUnction Failed :%s\n", err)
		os.Exit(1)
	}
	if err := myprinter(""); err != nil {
		if err == errorEmptyString {
			fmt.Printf("Custom Error error EmptyString :%s\n", err)
		} else {
			fmt.Printf("Printer Function Failed :%s\n", err)
		}
			os.Exit(1)
	}
}

