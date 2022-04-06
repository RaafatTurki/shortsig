package log

import (
	"fmt"
	"net"
)

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

var logTypeNames = map[int]string{
  ERR:    "ERROR",
  WARN:   "WARN",
  INFO:   "INFO",
  NET:    "NET",
  DEBUG:  "DEBUG",
}

var logTypeColors = map[int]string{
  ERR: red,
  WARN: yellow,
  INFO: green,
  NET: blue,
  DEBUG: magenta,
}

func LogMessageFoprmatter(logType int, msgFmtStr string, msgData... any) string {
  message := fmt.Sprintf(msgFmtStr, msgData...)

  logMessage := fmt.Sprintf(
    "%s %s %s %s\n",
    logTypeColors[logType], logTypeNames[logType], reset,
    message,
  )

  return logMessage
}

func LogMessageFoprmatterNoColor(logType int, msgFmtStr string, msgData... any) string {
  message := fmt.Sprintf(msgFmtStr, msgData...)

  logMessage := fmt.Sprintf(
    "%s %s\n",
    logTypeNames[logType],
    message,
  )

  return logMessage
}


func Log(logType int, msgFmtStr string, msgData... any) {
  fmt.Print(LogMessageFoprmatter(logType, msgFmtStr, msgData...))
}

func TCPLog(conn net.Conn, logType int, msgFmtStr string, msgData... any) {
  conn.Write([]byte(fmt.Sprintf("%s", LogMessageFoprmatterNoColor(logType, msgFmtStr, msgData...))))
}

func LogAndTCPLog(conn net.Conn, logType int, msgFmtStr string, msgData... any) {
  Log(logType, msgFmtStr, msgData...)
  TCPLog(conn, logType, msgFmtStr, msgData...)
}

func Assert(err error) {
  if err != nil {
    logMessage := fmt.Sprintf(
      "%s %s %s %s\n",
      logTypeColors[ERR], logTypeNames[ERR], reset,
      err,
    )

    fmt.Print(logMessage)
    panic("")
  }
}
