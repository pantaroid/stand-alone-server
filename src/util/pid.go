package util

import (
  "os"
  "io/ioutil"
  "path/filepath"
  "fmt"
)

type File struct {
  *os.File
  path string
}

func (f *File) Close() error {
  if err := f.File.Close(); err != nil {
    os.Remove(f.File.Name())
    return err
  }
  if err := os.Rename(f.Name(), f.path); err != nil {
    return err
  }
  return nil
}

func WritePidfile(path string) error {
  if err := os.MkdirAll(filepath.Dir(path), os.FileMode(0755)); err != nil {
    return err
  }
  f, err := ioutil.TempFile(filepath.Dir(path), filepath.Base(path))
  if err != nil {
    return err
  }
  if err = os.Chmod(f.Name(), os.FileMode(0644)); err != nil {
    f.Close()
    os.Remove(f.Name())
    return err
  }
  file := &File{File: f, path: path}
  if err != nil {
    return fmt.Errorf("error opening pidfile %s: %s", path, err)
  }
  defer file.Close()
  _, err = fmt.Fprintf(file, "%d", os.Getpid())
  if err != nil {
    return err
  }
  err = file.Close()
  if err != nil {
    return err
  }
  return nil
}

