package cmd

import (
  "bytes"
  "fmt"
  "log"
  "os/exec"
)

func launchForResult(exename string , args []string) (chan string, error) {
  out := make(chan string)

  exepath, err := exec.LookPath(exename)
  if err != nil {
    return nil, err
  }
  //fmt.Println(exepath)

  go func() {
    defer close(out)

    fmt.Println(exepath, args)
    cmd := exec.Command(exepath, args...)

    var stdout bytes.Buffer
    cmd.Stdout = &stdout

    //var stderr bytes.Buffer
    //cmd.Stderr = &stderr

    if err = cmd.Start(); err != nil {
      log.Panic(err)                          /* probably shouldn't exit */
    }

    if err = cmd.Wait(); err != nil {
      log.Panic(err)                          /* probably shouldn't exit */
    }

    out <- stdout.String()
  }()

  return out, err
}



