package main

import (
  "fmt"
  "runtime"
  "path"
  "os"
  "strconv"
  "github.com/kunterbunt/fileserver/server"
)

func main() {
  _, filename, _, ok := runtime.Caller(0)
  if !ok {
    panic("No caller information")
  }
  // Default port number.
  port := 8080
  // We need a directory to serve from.
  if len(os.Args) < 2 {
    fmt.Println("Provide path to serve files from.")
    return
  }
  dir := os.Args[1]
  // Optionally set the user-provided port, too.
  if len(os.Args) == 3 {
    var err error
    port, err = strconv.Atoi(os.Args[2])
    if err != nil {
      panic(err)
    }
  }
  // Instantiate server.
  server := server.NewServer(path.Dir(filename), port)
  // Fire it up.
  fmt.Printf("Serving \"" + dir + "\" from port %d.", port)
  server.ServeFromDirectory("/", dir)
  server.ListenAndServe()
}
