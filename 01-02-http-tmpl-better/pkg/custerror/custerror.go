package custerror

import (
	"fmt"
	"log"
)

func ErrPrint(err error) {
	if err != nil {
		fmt.Println("Error occured", err)
		return
	}
}

func FatalErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
