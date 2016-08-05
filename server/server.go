package server

import (
  "log"
  "strconv"	
  "net/http"
  "github.com/gorilla/mux"
)

type Server struct {
    router *mux.Router
    rootDirectory string
    port int
}

func NewServer(rootDirectory string, port int) *Server {
    var server Server
    server.router = mux.NewRouter()
    server.rootDirectory = rootDirectory
    server.port = port
    return &server
}

func (server Server) ServeFromDirectory(route string, directory string) {
  server.router.PathPrefix(route).Handler(http.StripPrefix(route, http.FileServer(http.Dir(directory))))
}

func (server Server) ListenAndServe() {
    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(server.port), server.router))
}
