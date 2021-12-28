package main

import (
  "html/template"
  "os"
)

var x = `
  <!DOCTYPE html>
  <html lang="en">
  <head>
  <meta charset="UTF-8">
  <title>Hello World</title>
  </head>
  <body>
    Hello {{.}}
  </body>
  </html>
`

func main() {
  t, err := template.New("hello").Parse(x)
  if err != nil {
    panic(err)
  }
  _ = t.Execute(os.Stdout, "<script>alert('world')</script>")
}
