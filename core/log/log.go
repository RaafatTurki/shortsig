package log

import (
	"fmt"
	"net"
)

const maxLogLevel = NET

const (
  red     = "\033[97;41m"
	green   = "\033[90;42m"
  yellow  = "\033[90;43m"
  blue    = "\033[97;44m"
  magenta = "\033[97;45m"
  cyan    = "\033[97;46m"
	white   = "\033[90;47m"
	reset   = "\033[0m"
)

const (
  ERR = iota
  WARN
  INFO
  NET
  DEBUG
)

var logLvlNames = map[int]string{
  ERR:    "ERROR",
  WARN:   "WARN",
  INFO:   "INFO",
  NET:    "NET",
  DEBUG:  "DEBUG",
}

var logLvlColors = map[int]string{
  ERR: red,
  WARN: yellow,
  INFO: green,
  NET: blue,
  DEBUG: magenta,
}

func MessageFormatter(logLvl int, color bool, msgFmtStr string, msgData... any) string {
  message := fmt.Sprintf(msgFmtStr, msgData...)
  var formattedMessage string

  if (color) {
    formattedMessage = fmt.Sprintf("%s %s %s %s\n", logLvlColors[logLvl], logLvlNames[logLvl], reset, message)
  } else {
    formattedMessage = fmt.Sprintf( "%s %s\n", logLvlNames[logLvl], message)
  }

  return formattedMessage
}

func PrintConsole(logLvl int, msgFmtStr string, msgData... any) {
  if logLvl > maxLogLevel { return }
  fmt.Print(MessageFormatter(logLvl, true, msgFmtStr, msgData...))
}

func PrintTCP(conn net.Conn, logLvl int, msgFmtStr string, msgData... any) {
  if logLvl > maxLogLevel { return }
  conn.Write([]byte(fmt.Sprintf("%s", MessageFormatter(logLvl, false, msgFmtStr, msgData...))))
}

func PrintConsoleAndTCP(conn net.Conn, logLvl int, msgFmtStr string, msgData... any) {
  PrintConsole(logLvl, msgFmtStr, msgData...)
  PrintTCP(conn, logLvl, msgFmtStr, msgData...)
}

func PrintErr(err error) {
  if err != nil {
    fmt.Print(MessageFormatter(ERR, true, "%s", err))
  }
}

func PanicErr(err error) {
  PrintErr(err)
  if err != nil {
    panic("")
  }
}
