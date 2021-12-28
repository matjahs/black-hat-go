package main

import (
  "context"
  "fmt"
  "log"
  "net/http"

  "github.com/gorilla/mux"
  "github.com/urfave/negroni"
)

type badAuth struct {
  Username string
  Password string
}

func (b *badAuth) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  username := r.URL.Query().Get("username")
  password := r.URL.Query().Get("password")
  if username != b.Username || password != b.Password {
    http.Error(w, "Unauthorized", http.StatusUnauthorized)
    return
  }
  ctx := context.WithValue(r.Context(), "username", username)
  r = r.WithContext(ctx)
  next(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {
  username := r.Context().Value("username")
  if _, err := fmt.Fprintf(w, "Hi %s\n", username.(string)); err != nil {
    log.Println(err)
  }
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/hello", hello).Methods("GET")
  n := negroni.Classic()

  n.Use(&badAuth{
    Username: "admin",
    Password: "password",
  })

  n.UseHandler(r)
  if err := http.ListenAndServe(":8000", n); err != nil {
    log.Fatalln(err)
  }
}
