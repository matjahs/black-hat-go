package main

import (
	"fmt"
	"io"
	"os"
)

type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (FooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	// input := make([]byte, 4096)

	// s, err := reader.Read(input)
	// if err != nil {
	// 	log.Fatalln("unable to read data")
	// }
	// fmt.Printf("read %d bytes from stdin\n", s)
	//
	// s, err = writer.Write(input)
	// if err != nil {
	// 	log.Fatalln("unable to write data")
	// }
	// fmt.Printf("wrote %d bytes to stdout", s)

	if _, err := io.Copy(&writer, &reader); err != nil {
		fmt.Println("unable to copy input to output")
	}
}
