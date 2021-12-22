package cmd

import (
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "runtime"
  "strings"
)

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

func KronkEmbed0(pkgName, varName, filename string) error {
  dirname := filepath.Join(MyDirName(), "defassets")
  return KronkEmbed(pkgName, varName, filename, dirname)
}

func KronkEmbed(pkgName, varName, filename string, dirs ...string) error {
  //verbose0(fmt.Sprintf("pkg: %s, var: %s, out: %s\n", pkgName, varName, filename))
  w, err := os.Create(filename)
  if err != nil {
    return err
  }
  defer w.Close()

  fmt.Fprintf(w, "package %s\n", pkgName)
  fmt.Fprintf(w, "%s", embedHelperZero())

  // Generate the header that shows files
  fmt.Fprintf(w, "// -----------------------------\n// File list:\n//\n")
  for _, dir := range dirs {
    filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
      if ! info.IsDir() {
        relative := filepath.ToSlash(strings.TrimPrefix(path, dir))
        fmt.Fprintf(w, "// %s\n", relative)
      }
      return nil
    })
  }
  fmt.Fprintf(w, "// -----------------------------\n")

  fmt.Fprintf(w, "\n%s", embedHelperOne())
  fmt.Fprintf(w, "var %s = &fs{}\n", varName)
  fmt.Fprintf(w, "%s", embedHelperTwo())
  fmt.Fprintf(w, "%s\n", "func init() {")
  defer fmt.Fprintf(w, "}\n")

  // Generate all the data
  for _, dir := range dirs {
    filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
      if info.IsDir() {
        Vvverbose(fmt.Sprintf("Not asseting (ISDIR) %s\n", path))
        return nil
      }
      b, err := ioutil.ReadFile(path)
      if err != nil {
        return err
      }
      if len(b) == 0 {
        Vvverbose(fmt.Sprintf("Skipping zero-length file: %s\n", path))
        return nil
      }
      Verbose(fmt.Sprintf("asseting %s (len: %d)\n", path, len(b)))

      relative := filepath.ToSlash(strings.TrimPrefix(path, dir))
      fmt.Fprintf(w, "  assets[%q] = []byte{\n    ", relative)
      defer fmt.Fprintf(w, "%s\n", "}")

      fmt.Fprintf(w, "0x%02x", b[0])
      for i := 1; i < len(b); i++ {
        if i % 16 == 0 {
          fmt.Fprintf(w, ",\n    0x%02x", b[i])
        } else {
          fmt.Fprintf(w, ", 0x%02x", b[i])
        }
      }

      return nil
    })
  }

  return nil
}

func embedHelperZero() string {
  return `

// Code auto-generated. DO NOT EDIT.

`
}

func embedHelperOne() string {
  return `

import (
  "bytes"
  "errors"
  "net/http"
  "os"
  "time"
)

var assets = map[string][]byte{}

`
}

func embedHelperTwo() string {
  return `

type fs struct {}

func (fs *fs) Open(name string) (http.File, error) {
  if name == "/" {
    return fs, nil;
  }
  b, ok := assets[name]
  if !ok {
    return nil, os.ErrNotExist
  }
  return &file{name: name, size: len(b), Reader: bytes.NewReader(b)}, nil
}

func (fs *fs) Close() error { return nil }
func (fs *fs) Read(p []byte) (int, error) { return 0, nil }
func (fs *fs) Seek(offset int64, whence int) (int64, error) { return 0, nil }
func (fs *fs) Stat() (os.FileInfo, error) { return fs, nil }
func (fs *fs) Name() string { return "/" }
func (fs *fs) Size() int64 { return 0 }
func (fs *fs) Mode() os.FileMode { return 0755}
func (fs *fs) ModTime() time.Time{ return time.Time{} }
func (fs *fs) IsDir() bool { return true }
func (fs *fs) Sys() interface{} { return nil }

func (fs *fs) Readdir(count int) ([]os.FileInfo, error) {
  files := []os.FileInfo{}
  for name, data := range assets {
    files = append(files, &file{name: name, size: len(data), Reader: bytes.NewReader(data)})
  }
  return files, nil
}

type file struct {
  name string
  size int
  *bytes.Reader
}

func (f *file) Close() error { return nil }
func (f *file) Readdir(count int) ([]os.FileInfo, error) { return nil, errors.New("not supported") }
func (f *file) Stat() (os.FileInfo, error) { return f, nil }
func (f *file) Name() string { return f.name }
func (f *file) Size() int64 { return int64(f.size) }
func (f *file) Mode() os.FileMode { return 0644 }
func (f *file) ModTime() time.Time{ return time.Time{} }
func (f *file) IsDir() bool { return false }
func (f *file) Sys() interface{} { return nil }

`

}

func file_line() string {
  _, fileName, fileLine, ok := runtime.Caller(1)
  var s string
  if ok {
    s = fmt.Sprintf("%s:%d", fileName, fileLine)
  } else {
    s = ""
  }
  return s
}


