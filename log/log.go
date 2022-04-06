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

func LogMessageFoprmatter(logLvl int, msgFmtStr string, msgData... any) string {
  message := fmt.Sprintf(msgFmtStr, msgData...)

  logMessage := fmt.Sprintf(
    "%s %s %s %s\n",
    logLvlColors[logLvl], logLvlNames[logLvl], reset,
    message,
  )

  return logMessage
}

func LogMessageFoprmatterNoColor(logLvl int, msgFmtStr string, msgData... any) string {
  message := fmt.Sprintf(msgFmtStr, msgData...)

  logMessage := fmt.Sprintf(
    "%s %s\n",
    logLvlNames[logLvl],
    message,
  )

  return logMessage
}


func Log(logLvl int, msgFmtStr string, msgData... any) {
  if logLvl > maxLogLevel { return }
  fmt.Print(LogMessageFoprmatter(logLvl, msgFmtStr, msgData...))
}

func TCPLog(conn net.Conn, logLvl int, msgFmtStr string, msgData... any) {
  if logLvl > maxLogLevel { return }
  conn.Write([]byte(fmt.Sprintf("%s", LogMessageFoprmatterNoColor(logLvl, msgFmtStr, msgData...))))
}

func LogAndTCPLog(conn net.Conn, logLvl int, msgFmtStr string, msgData... any) {
  Log(logLvl, msgFmtStr, msgData...)
  TCPLog(conn, logLvl, msgFmtStr, msgData...)
}

func Assert(err error) {
  if err != nil {
    logMessage := fmt.Sprintf(
      "%s %s %s %s\n",
      logLvlColors[ERR], logLvlNames[ERR], reset,
      err,
    )

    fmt.Print(logMessage)
    panic("")
  }
}
