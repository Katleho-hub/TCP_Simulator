package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin) // reading user inputs
		fmt.Println(">> ")
		text, _ := reader.ReadString('\n') // reading user inputs
		fmt.Fprintf(c, text+"\n")          // sending user inputs to server

		message, _ := bufio.NewReader(c).ReadString('\n') // reading server response
		fmt.Print("->: " + message)                       //console logging server response

		if strings.TrimSpace(string(text)) == "STOP" { // if user input is "STOP"
			fmt.Println("TCP client exiting...")
			return
		}

	}
}
