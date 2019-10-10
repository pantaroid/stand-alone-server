package main

import (
  "flag"
  "fmt"
  "io/ioutil"
  "net"
  "net/http"
  "os"
  "os/exec"
  "os/signal"
  "./util"
)
const (
  SOCKET = "server.sock"
  PID    = "server.pid"
)
func main() {
  if 1 == len(os.Args) {
    cmd := exec.Command(os.Args[0], "start")
    cmd.Start()
  } else {
    c, _ := os.Getwd()
    flag.Parse()
    if 1 == flag.NArg() {
      pid := fmt.Sprintf("%s/%s", c, PID)
      switch flag.Arg(0) {
      case "start":
        listener, err := start()
        if err != nil {
          util.Errorf("%q", err)
        } else {
          defer stop(listener)
          err = util.WritePidfile(pid)
          if err == nil {
            quit := make(chan os.Signal)
            signal.Notify(quit, os.Interrupt)
            <-quit
            util.Info("receive quit signal.")
            stop(listener)
          } else {
            util.Errorf("cannot create pid file: %s", err)
            stop(listener)
          }
        }
      case "stop":
        util.Info("server stopping")
        _, err := os.Stat(pid)
        if err == nil {
          bytes, err := ioutil.ReadFile(pid)
          util.Infof("send kill signal to %s", string(bytes))
          out, err := exec.Command("kill", "-2", string(bytes)).Output()
          util.Infof("%q", out)
          if err != nil {
            util.Errorf("kill pid: %s", out)
          }
        } else {
          util.Errorf("stop pid error: %s", err)
        }
      }
    }
  }
}

func start() (net.Listener, error) {
  util.Info("start server")
  listener, err := net.Listen("unix", SOCKET)
  if err != nil {
    util.Errorf("listen (unix socket): %s", err)
    listener.Close()
    return nil, err
  }
  // TODO: Add your functions
  // Ex: http.HandleFunc(action.New().GetHandleInfo())
  go func() {
    util.Info("listener set.")
    os.Chmod(SOCKET, 0777)
    http.Serve(listener, nil)
  }()
  return listener, nil
}

func stop(listener net.Listener) {
  _, err := os.Stat(SOCKET)
  if err == nil {
    util.Info("stop server")
    os.Remove(SOCKET)
    os.Remove(PID)
    listener.Close()
  } else {
    util.Errorf("stop error: %s", err)
  }
}
