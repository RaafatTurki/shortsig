package main

import (
	"bufio"
	"fmt"
	"net"
	"shortsig/config"
	. "shortsig/log"
	"shortsig/service"
	"strings"
)

var conf config.Config

func main() {
  conf = config.ParseConfigFile("config.toml")

  // spinning tcp server
  listener, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
  Assert(err)
  Log(INFO, "listening on port %d", conf.Port)

  // listening to incoming tcp connections
  for {
    conn, err := listener.Accept()
    if err != nil {
      Log(WARN, "connection accept error %s", err)
      continue
    }

    // whitelising
    // connIP := strings.Split(conn.RemoteAddr().String(), ":")[0]
    //
    // is_whitelisted := false
    // for _, ip := range conf.Whitelist {
    //   if ip == connIP { is_whitelisted = true }
    // }
    //
    // if (!is_whitelisted) {
    //   log.Printf("IP %s not whitelisted", connIP)
    //   continue
    // }

    // accepting connections and handling them in a new thread
    LogAndTCPLog(conn, NET, "connection accepted %s", conn.RemoteAddr().String())
    go handleConnection(conn)
  }
}

func handleConnection(conn net.Conn) {
  defer conn.Close()
  s := bufio.NewScanner(conn)

  for s.Scan() {
    // Log(NET, "client:(%s): Incoming Data", conn.RemoteAddr())
    data := s.Text()

    if !handleTCPCmd(data, conn) {
      break
    }
  }
}

func handleTCPCmd(TCPData string, conn net.Conn) bool {
  TCPDataArr := strings.Split(TCPData, " ")
  TCPCmd := TCPDataArr[0]
  TCPCmdArgs := TCPDataArr[1:]

  switch TCPCmd {
    case "":
      LogAndTCPLog(conn, NET, "client(%s): Empty TCP Command", conn.RemoteAddr())
    case "exit":
      LogAndTCPLog(conn, NET, "client(%s): Disconnected", conn.RemoteAddr())
    case "cmd":
      LogAndTCPLog(conn, NET, "client(%s): Executing TCP Command %s", conn.RemoteAddr(), TCPCmdArgs)
      service.ExecCommand(conn, TCPCmdArgs, conf.Routines)
      LogAndTCPLog(conn, NET, "client(%s): Executed TCP Command %s", conn.RemoteAddr(), TCPCmdArgs)
    default:
      LogAndTCPLog(conn, NET, "client(%s): Invalid TCP Command", conn.RemoteAddr().String())
  }
  return true
}
