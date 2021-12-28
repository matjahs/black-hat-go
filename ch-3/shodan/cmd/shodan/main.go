package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/matjahs/black-hat-go/ch-3/shodan/shodan"
)

func main() {
  if len(os.Args) != 2 {
    log.Fatalln("usage: shodan <query>")
  }

  if err := godotenv.Load("/Users/matjah/Code/gitlab.com/black-hat-go/ch-3/shodan/.env"); err != nil {
    log.Fatalln(err)
  }
  apiKey := os.Getenv("API_TOKEN")
  s := shodan.New(apiKey)
  info, err := s.APIInfo()
  if err != nil {
    log.Panicln(err)
  }
  fmt.Printf(
    "Query Credits: %d\nScan Credits: %d\n\n",
    info.QueryCredits,
    info.ScanCredits,
  )

  hostSearch, err := s.HostSearch(os.Args[1])
  if err != nil {
    log.Panicln(err)
  }
  for _, host := range hostSearch.Matches {
    fmt.Printf("%18s%8d\n", host.IPString, host.Port)
  }

}
