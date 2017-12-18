package main

import(
  "fmt"
  "flag"
  "server"
)

func main() {
  serverPtr := flag.Bool("server", false, "Server")
  clientPtr := flag.Bool("client", false, "Client")
  hostPtr := flag.String("host", "", "Server Host")
  portPtr := flag.String("port", "", "Server Port")
  flag.Parse()

  if *serverPtr && !*clientPtr {
    fmt.Println("Starting as server")
    srv := server.Server{}
    srv.Start()
  } else if *hostPtr != "" && *portPtr != "" {
    fmt.Println("Starting as client")
    fmt.Printf("Host : %v Port : %v\n", *hostPtr, *portPtr)
  } else {
    flag.Usage()
  }
}
