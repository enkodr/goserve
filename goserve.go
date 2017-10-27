package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.Handle("/", http.FileServer(http.Dir("./")))
  fmt.Printf("Listening on <http://127.0.0.1:8000>")
  http.ListenAndServe(":8000", nil)
}
