package client

import(
  "bufio"
  "net"
  "math/rand"
)

func main() {
  connection, _ := net.Dial("tcp", "127.0.0.1:8081")

  cli := client.NewClient("name1", connection)
  cli.StartListener()
}

type Client struct {
  ID        string
  Name      string
  Conn      net.Conn
  Reader    *bufio.Reader
  Writer    *bufio.Writer
}

func NewClient(name string, conn net.Conn) *Client {
  id := rand.Int()
  fmt.Println("Client Id : ", id)
  client := &Client{
    ID:     id,
    Name:   name,
    Conn:   conn,
    Reader: bufio.NewReader(c),
    Writer: bufio.NewWriter(c),
  }

  return client
}

func (client *Client) SendMessage(msg string) {
  client.Writer.WriteString(msg)
  client.Writer.Flush()
}

func (client *Client) StartListener() {
  for {
    message, _ := client.Reader.ReadString("\n")
    fmt.Println("Message Received : ", message)
  }
}
