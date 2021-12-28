package main

import (
  "fmt"
  "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
  _, _ = fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}

func main() {
  http.HandleFunc("/", hello)
  _ = http.ListenAndServe(":8000", nil)
}
