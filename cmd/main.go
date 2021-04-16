package main

import (
	"flag"
	"fmt"
    "log"
	"net/http"
)

const (
    defaultPort = "3000"
    defaultDir = "./"
)

var (
   port string
   dir string
)

func init() {
    flag.StringVar(&port, "port", defaultPort, "Port number to listen on.")
    flag.StringVar(&dir, "dir", defaultDir, "Directory path to serve.")
    flag.Parse()
}

func main() {
    http.Handle("/", http.FileServer(http.Dir(dir)))
    fmt.Printf("Listening on port %s...", port)
log.Fatal(http.ListenAndServe(":"+port, nil))
}
