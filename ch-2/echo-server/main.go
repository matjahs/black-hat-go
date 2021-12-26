package main

import (
  "io"
  "log"
  "net"
)

func echo(conn net.Conn) {
  defer conn.Close()

  b := make([]byte, 512)
  for {
    size, err := conn.Read(b[0:])
    if err != io.EOF {
      log.Println("client disconnected")
      break
    }
    if err != nil {
      log.Println("unexpected error")
      break
    }
    log.Printf("received %d bytes: %s\n", size, string(b))

    log.Println("writing data")
    if _, err := conn.Write(b[0:size]); err != nil {
      log.Fatalln("unable to write data")
    }
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
