package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

type Message struct {
	Name    string
	Content string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Welcome %s! You've joined the Chat.\n", name)
	for {
		fmt.Print("Enter message (or 'exit' to quit): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			break
		}

		var chatHistory []Message
		msg := Message{Name: name, Content: text}

		err = client.Call("ChatServer.SendMessage", msg, &chatHistory)
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}

		fmt.Println("\n--- Chat History ---")
		for _, m := range chatHistory {
			fmt.Printf("%s: %s\n", m.Name, m.Content)
		}
		fmt.Println()
	}
}