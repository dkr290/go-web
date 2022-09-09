package custerror

import "fmt"

func ErrPrint(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
