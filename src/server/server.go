package server

import(
  "fmt"
  "net"
  "bufio"
)

type Server struct {

}

func (server *Server) Start() {
  fmt.Println("Starting server...")

  listener, err := net.Listen("tcp", ":8081")
  if err != nil {
    fmt.Println("Could not start server. ")
    panic(err)
  }

  for {
    connection, err := listener.Accept()
    if err != nil {
      panic(err)
    }
  }
}
