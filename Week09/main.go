package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)


type Message struct {
	message string
	ch      chan []byte
	wg      *sync.WaitGroup
}

func (m *Message) NewServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("start server failed!")
	}
	fmt.Println("start server...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		m.wg.Add(2)
		go m.handleRead(conn)
		go m.handleWrite(conn)
		m.wg.Wait()
	}
}

func (m *Message) handleWrite(conn net.Conn) {
	defer m.wg.Done()
	for i := 10; i > 0; i-- {
		line, _ := <- m.ch
		_, err := conn.Write(line)
		_, err = conn.Write([]byte("\n"))
		if err != nil {
			fmt.Println("Error to send message because of ", err.Error())
			break
		}
	}
}
func (m *Message) handleRead(conn net.Conn) {
	defer m.wg.Done()
	reader := bufio.NewReader(conn)
	for i := 1; i <= 10; i++ {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Print("Error to read message because of ", err)
			return
		}
		m.ch <- line
	}
}

func main() {
	ch := make(chan []byte, 2)
	wg := sync.WaitGroup{}
	m := Message{
		message: "",
		ch: ch,
		wg: &wg,
	}
	m.NewServer()
}
