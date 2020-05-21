package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
)

const (
	PROTOCOL = "tcp"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Arugemnts invalides host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial(PROTOCOL, CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "QUIT" {
				return
		}
	}
}