package util

import (
  "log"
  "os"
  "sync"
)

var (
  stdout = log.New(os.Stdout, "", log.LstdFlags)
  stderr = log.New(os.Stderr, "", log.LstdFlags)
  logger *log.Logger
  mutex  sync.Mutex
)

func SetLogFile(file *os.File) {
  logger = log.New(file, "", log.LstdFlags)
}
func Info(message string) {
  mutex.Lock()
  if logger != nil {
    logger.Println(message)
  } else {
    stdout.Println(message)
  }
  mutex.Unlock()
}
func Infof(format string, args ...interface{}) {
  mutex.Lock()
  if logger != nil {
    logger.Printf(format + "\n", args...)
  } else {
    stdout.Printf(format + "\n", args...)
  }
  mutex.Unlock()
}
func Error(message string) {
  mutex.Lock()
  if logger != nil {
    logger.Println(message)
  } else {
    stderr.Println(message)
  }
  mutex.Unlock()
}
func Errorf(format string, args ...interface{}) {
  mutex.Lock()
  if logger != nil {
    logger.Printf(format + "\n", args...)
  } else {
    stderr.Printf(format + "\n", args...)
  }
  mutex.Unlock()
}

