package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("aks")
	if err != nil {
		fmt.Print("There has been an error!: ", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
