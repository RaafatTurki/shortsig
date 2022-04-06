package service

import (
	"fmt"
	"net"
	"os"
	"os/exec"
)

func ExecCommand(conn net.Conn, payload []string, cmdsTable map[string]string) {
  cmdKey := payload[0]
  cmd, ok := cmdsTable[cmdKey]
  if !ok { fmt.Printf("client(%s): Invalid Command Payload %s", conn.RemoteAddr(), cmdKey) }

  cmd_exec := exec.Command("sh", "-c", cmd)

  cmd_exec.Stdin = os.Stdin;
  cmd_exec.Stdout = os.Stdout;
  cmd_exec.Stderr = os.Stderr;
  err := cmd_exec.Run()
  if err != nil { fmt.Printf("%v\n", err) }
}

// THIS IS A GIANT SECURITY HAZARD AND SHOULD NEVER BE USED OUTSIDE TESTING
// execArbitraryCommand(conn net.Conn, payload []string, cmdsTable map[string]string) {}
