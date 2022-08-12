package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT) // creating tcp server
	if err != nil {
		fmt.Println(err)
		return
	}

	defer l.Close()

	c, err := l.Accept() // on accept we allow communication with client
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n') // reading client inputs
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(netData)) == "STOP" { // if client input is "STOP"
			fmt.Println("TCP client exiting...")
			return
		}

		fmt.Print("->: " + string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n" //console logging server response

		c.Write([]byte(myTime))

	}
}
