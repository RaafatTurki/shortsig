package service

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"shortsig/config"
)

func ExecRoutine(conn net.Conn, payload []string, routines map[string]config.Routine) {
  routineName := payload[0]
  routine, ok := routines[routineName]
  if !ok { fmt.Printf("%s | Invalid Routine %s", conn.RemoteAddr(), routineName) }

  var cmd *exec.Cmd

  platform := runtime.GOOS
  switch platform {
  case "windows":
    cmd = exec.Command("cmd", "/C", routine.Windows)
  case "darwin":
    cmd = exec.Command("sh", "-c", routine.Darwin)
  case "linux":
    cmd = exec.Command("sh", "-c", routine.Linux)
  default:
    cmd = exec.Command("sh", "-c", routine.Other)
  }

  cmd.Stdin = os.Stdin;
  cmd.Stdout = os.Stdout;
  cmd.Stderr = os.Stderr;
  err := cmd.Run()
  if err != nil { fmt.Printf("%v\n", err) }
}

// THIS IS A GIANT SECURITY HAZARD AND SHOULD NEVER BE USED OUTSIDE TESTING
// execArbitraryCommand(conn net.Conn, payload []string, cmdsTable map[string]string) {}
