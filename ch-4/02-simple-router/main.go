package main

import (
  "fmt"
  "net/http"
)

type router struct{}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  switch req.URL.Path {
  case "/a":
    _, _ = fmt.Fprintf(w, "Executing /a")
  case "/b":
    _, _ = fmt.Fprintf(w, "Executing /b")
  case "/c":
    _, _ = fmt.Fprintf(w, "Executing /c")
  default:
    http.Error(w, "404 Not Found", http.StatusNotFound)
  }
}

func main() {
  var r router
  _ = http.ListenAndServe(":8000", &r)
}
