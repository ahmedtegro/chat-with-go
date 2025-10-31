package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Message struct to hold chat messages
type Message struct {
	Name    string
	Content string
}

// ChatServer struct to store all messages
type ChatServer struct {
	history []Message
}

// SendMessage adds new message to history and returns all messages
func (c *ChatServer) SendMessage(msg Message, reply *[]Message) error {
	c.history = append(c.history, msg)
	*reply = c.history
	return nil
}

func main() {
	server := new(ChatServer)
	rpc.Register(server)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Chat server running on port 1234...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}