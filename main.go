package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"shortsig/service"
	"strings"
)

var conf Config

func main() {
  conf = ParseConfigFile("config.toml")

  // spinning tcp server
  listener, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
  if err != nil { panic(err) }
  log.Printf("listening on port %d", conf.Port)

  // listening to incoming tcp connections
  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Printf("Accept Error %s", err)
      continue
    }

    connIP := strings.Split(conn.RemoteAddr().String(), ":")[0]

    is_whitelisted := false
    for _, ip := range conf.Whitelist {
      if ip == connIP { is_whitelisted = true }
    }

    if (!is_whitelisted) {
      log.Printf("IP %s not whitelisted", connIP)
      continue
    }

    // accepting connections and handling them in a new thread
    log.Printf("Accepted %s", conn.RemoteAddr())
    go handleConnection(conn)
  }
}

func handleConnection(conn net.Conn) {
  defer conn.Close()
  conn.Write([]byte("SERVER: connected to server\n"))
  s := bufio.NewScanner(conn)

  for s.Scan() {
    // log.Printf("client:(%s): Incoming Data", conn.RemoteAddr())
    data := s.Text()

    if !handleTCPCmd(data, conn) {
      break
    }
  }
}

func handleTCPCmd(TCPdata string, conn net.Conn) bool {
  TCPdataArr := strings.Split(TCPdata, " ")
  TCPcmd := TCPdataArr[0]

  switch TCPcmd {
    case "":
      log.Printf("client(%s): Empty Command", conn.RemoteAddr())
      conn.Write([]byte("SERVER: empty command\n"))
      return true
    case "exit":
      log.Printf("client(%s): Disconnected", conn.RemoteAddr())
      conn.Write([]byte("SERVER: disconnected\n"))
      return false
    case "exec":
      // service.SpawnSubprocessFromCommand("ls")
      service.ExecCommand(conn, TCPdataArr[1:], conf.Cmds)
      // get(str[1:], conn)
      return true
    default:
      log.Printf("client(%s): Invalid Command", conn.RemoteAddr().String())
      conn.Write([]byte("SERVER: invalid payload\n"))
      return true
  }
}
