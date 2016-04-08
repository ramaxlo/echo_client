package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/websocket"
	"os"
)

func main() {
	ws, err := websocket.Dial("ws://localhost:8080/echo", "", "http://localhost/")
	if err != nil {
		fmt.Println("Fail to dial")
		return
	}

	/*
		_, err = ws.Write([]byte("echo test\n"))
		if err != nil {
			fmt.Println("Fail to write")
			return
		}
	*/

	var n int
	br := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Input: ")
		msg, err := br.ReadString('\n')
		if err != nil {
			fmt.Println("Fail to read")
			return
		}

		if msg == "\n" {
			continue
		}

		if msg == "q\n" || msg == "Q\n" {
			break
		}

		_, err = ws.Write([]byte(msg))
		if err != nil {
			fmt.Println("Fail to write")
			return
		}

		var recv = make([]byte, 512)
		n, err = ws.Read(recv)
		if err != nil {
			fmt.Println("Fail to read")
			return
		}

		fmt.Printf("Received: %s", string(recv[:n]))
	}

}
