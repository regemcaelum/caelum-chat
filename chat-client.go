package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func connectToServer(server string) net.Conn {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		// TODO Error handling
		log.Fatal(err)
	}
	return conn
}

func printPage(conn net.Conn) {
	var inString []byte
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	retCount, readErr := conn.Read(inString)
	if err != nil || readErr != nil {
		// TODO Error handling
		if readErr != nil {
			err = readErr
		}
		log.Fatal(err)
	}
	fmt.Print(status)
	fmt.Print(retCount)
	fmt.Print(string(inString))
}

func readStdIn() string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Name: ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			return text
		}
	}
}

func main() {
	conn := connectToServer("golang.org:80")
	printPage(conn)
	conn.Close()
}
