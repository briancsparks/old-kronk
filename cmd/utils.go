package cmd

import (
  "log"
  "path"
  "path/filepath"
  "runtime"
  "strings"
)

func checkU(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func q(s string) string {
  return "\"" + s + "\""
}

func qq(ss []string) []string {
  r := make([]string, 0)
  for _, s := range ss {
    r = append(r, q(s))
  }
  return r
}

func smap(ss []string, fn func(string) string) (res []string) {
  res = make([]string, 0)
  for _, s := range ss {
    res = append(res, fn(s))
  }
  return
}

func keyMirror(ss []string) (res map[string]string) {
  res = make(map[string]string)
  for _, s := range ss {
    res[s] = s
  }
  return
}

func keyMirror1(s string) (res map[string]string) {
  res = make(map[string]string)
  res[s] = s
  return
}

func keys(m map[string]string) (res []string) {
  res = make([]string, 0)
  for k, _ := range m {
    res = append(res, k)
  }
  return
}

func slc(s string) []string {
  return []string{s}
}

func slc0() []string {
  return []string{}
}

//dirsIn := smap(shortDirsIn, func(s string) string { return filepath.Join(dirname, s) })
//dirsIn := smap(shortDirsIn, prependPath(dirname))

func prependPath(pre string) func(s string) string {
  return func(s string) string {
    return filepath.Join(pre, s)
  }
}

func appendPath(suffix string) func(s string) string {
  return func(s string) string {
    return filepath.Join(s, suffix)
  }
}

func cygpath(s string) string {
  if len(s) < 2 || s[1] != ':' {
    return s
  }

  drive := strings.ToLower(s[0:1])
  path := strings.Replace("/cygdrive/" + drive + s[2:], "\\", "/", -1)
  return path
}

func MyFileName() string {
  _, filename, _, ok := runtime.Caller(1)
  if ok {
    return filename
  }
  return ""
}

func MyDirName() string {
  _, filename, _, ok := runtime.Caller(1)
  if ok {
    return path.Dir(filename)
  }
  return ""
}

