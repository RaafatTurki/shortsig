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
  // parse config
  conf = config.ParseConfigs()

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
    LogAndTCPLog(conn, NET, "%s | connection accepted", conn.RemoteAddr())
    go handleConnection(conn)
  }
}

func handleConnection(conn net.Conn) {
  defer conn.Close()
  s := bufio.NewScanner(conn)

  for s.Scan() {
    Log(DEBUG, "%s | Incoming Data", conn.RemoteAddr())
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
    LogAndTCPLog(conn, NET, "%s | Empty TCP Command", conn.RemoteAddr())
  case "exit":
    LogAndTCPLog(conn, NET, "%s | Disconnected", conn.RemoteAddr())
  case "exec":
    LogAndTCPLog(conn, NET, "%s | Executing TCP Command %s", conn.RemoteAddr(), TCPCmdArgs)
    service.ExecRoutine(conn, TCPCmdArgs, conf.Routines)
    LogAndTCPLog(conn, NET, "%s | Executed TCP Command %s", conn.RemoteAddr(), TCPCmdArgs)
  default:
    LogAndTCPLog(conn, NET, "%s | Invalid TCP Command", conn.RemoteAddr().String())
  }
  return true
}
