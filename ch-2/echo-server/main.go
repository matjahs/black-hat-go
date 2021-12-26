package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func handleEchoIO(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 512)

	size, err := conn.Read(b[0:])
	if err == io.EOF {
		log.Println("client disconnected")
		return
	}
	if err != nil {
		log.Println("unexpected error")
		return
	}
	log.Printf("received %d bytes: %s\n", size, string(b))

	log.Println("writing data")
	if _, err := conn.Write(b[0:size]); err != nil {
		log.Fatalln("unable to write data")
	}
}

func handleEchoBuffered(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("unable to read data")
		return
	}
	log.Printf("received %d bytes: %s\n", len(s), s)

	log.Println("writing data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("unable to write data")
	}
	_ = writer.Flush()
}

func handleEchoCopy(conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("unable to read/write data")
	}
}

func echo(conn net.Conn) {
	for {
		handleEchoBuffered(conn)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("unable to bind to port")
	}
	log.Println("listing on 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()
		log.Println("received connection")
		if err != nil {
			log.Fatalln("unable to accept connection")
		}
		go echo(conn)
	}
}
