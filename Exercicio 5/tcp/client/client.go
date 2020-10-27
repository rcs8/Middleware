package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	crhTCP, err := newCRHTCP()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type a message to server:")
	text, _ := reader.ReadString('\n')
	var response = string(crhTCP.SendReceive(text))
	fmt.Println("Server response: ", response)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	} else if err == io.EOF {
		log.Println("Connection closed.")
	}
}
