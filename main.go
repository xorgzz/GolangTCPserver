package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		// if "STOP" writen on client side the connection with client is ended
		if temp == "STOP" {
			break
		} else {
			inpt := string(netData)
			fmt.Print(inpt)

			// response to client of server
			out := "hello " + strings.Split(inpt, "\n")[0] + "\n"
			c.Write([]byte(string(out)))
		}

	}
	c.Close()
}

func main() {

	PORT := ":2137"
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)

	}
}
