package main

import (
	"io"
	"net"
	"testing"
)

func Test_echo(t *testing.T) {
	tests := []struct {
		name string
		fn   func(conn net.Conn)
		in   string
	}{
		{"handle using io", handleEchoIO, "test\n"},
		{"handle using bufio", handleEchoBuffered, "test\n"},
		{"handle using io.copy", handleEchoCopy, "test\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()

			go func() {
				handleEchoBuffered(server)
				server.Close()
			}()

			_, err := client.Write([]byte(tt.in))
			if err != nil {
				t.Fatalf("failed to write: %s", err)
			}

			in, err := io.ReadAll(client)
			if err != nil {
				t.Fatalf("failed to read: %s", err)
			}

			client.Close()

			if string(in) != tt.in {
				t.Fatalf("expected `test`, got %s", in)
			}
		})
	}
}
