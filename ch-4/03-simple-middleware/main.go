package main

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

// type logger struct {
//   Inner http.Handler
// }
//
// func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//   log.Println("start")
//   l.Inner.ServeHTTP(w, r)
//   log.Println("finish")
// }
//
// func hello(w http.ResponseWriter, _ *http.Request) {
//   fmt.Println(w, "Hello\n")
// }

func main() {
  // f := http.HandlerFunc(hello)
  // l := logger{Inner: f}
  // _ = http.ListenAndServe(":8000", &l)

  r := mux.NewRouter()

  r.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
    _, _ = fmt.Fprint(w, "hi foo")
  }).Methods("GET").Host("www.foo.com")

  r.HandleFunc("/users/{user}", func(w http.ResponseWriter, r *http.Request) {
    user := mux.Vars(r)["user"]
    _, _ = fmt.Fprintf(w, "hi %s\n", user)
  }).Methods("GET")

  _ = http.ListenAndServe(":8000", r)
}
