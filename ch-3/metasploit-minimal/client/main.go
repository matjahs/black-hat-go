package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/matjahs/black-hat-go/ch-3/metasploit-minimal/rpc"
)

func main() {
  _ = godotenv.Load()
  host := os.Getenv("MSF_HOST")
  pass := os.Getenv("MSF_PASS")
  user := "msf"

  if host == "" || pass == "" {
    log.Fatalln("Missing required environment variable MSF_HOST or MSF_PASS")
  }

  msf, err := rpc.New(host, user, pass)
  if err != nil {
    log.Panicln(err)
  }
  defer msf.Logout()

  sessions, err := msf.SessionList()
  if err != nil {
    log.Panicln(err)
  }
  fmt.Println("Sessions:")
  for _, session := range sessions {
    fmt.Printf("%5d  %s\n", session.ID, session.Info)
  }
}
