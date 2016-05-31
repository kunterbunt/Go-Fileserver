package main

import (
  "os"
	"fmt"
  "net/http"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Provide path to serve files from.")
    return
  }
  dir := os.Args[1]
  fileserver := http.FileServer(http.Dir(dir))
  http.Handle("/", fileserver)
  http.ListenAndServe(":3000", nil)
}
