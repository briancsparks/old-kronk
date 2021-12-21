package cmd

import (
  "bytes"
  "fmt"
  "log"
  "os/exec"
  "strings"
  "syscall"
)

func launch4Result(exename string , args []string) (chan string, error) {
  return launchForResult(exename, args, "", "")

  //out := make(chan string)
  //
  //exepath, err := exec.LookPath(exename)
  //if err != nil {
  //  return nil, err
  //}
  ////fmt.Println(exepath)
  //
  //go func() {
  //  defer close(out)
  //
  //  cmd := exec.Command(exepath, args...)
  //
  //  var stdout bytes.Buffer
  //  cmd.Stdout = &stdout
  //
  //  //var stderr bytes.Buffer
  //  //cmd.Stderr = &stderr
  //
  //  verbose0(fmt.Sprintf("  ----- launch: %s, %v\n", exepath, args))
  //  //fmt.Println(exepath, args)
  //  if err = cmd.Start(); err != nil {
  //    log.Panic(err)                          /* probably shouldn't exit */
  //  }
  //
  //  verbose0(fmt.Sprintf("  ----- launch2: %s, %v\n", exepath, args))
  //  if err = cmd.Wait(); err != nil {
  //    log.Panic(err)                          /* probably shouldn't exit */
  //  }
  //
  //  verbose0(fmt.Sprintf("  ----- launch3: %s, %v\n        %s\n", exepath, args, stdout))
  //  out <- stdout.String()
  //}()
  //
  //return out, err
}


func launchForResult(exename string, args []string, cwd string, deffault string) (chan string, error) {
  out := make(chan string)

  exepath, err := exec.LookPath(exename)
  if err != nil {
    return nil, err
  }
  //fmt.Println(exepath)

  go func() {
    defer close(out)

    res := deffault

    cmd := exec.Command(exepath, args...)

    if len(cwd) > 0 {
      cmd.Dir = cwd
    }

    var stdout bytes.Buffer
    cmd.Stdout = &stdout

    //var stderr bytes.Buffer
    //cmd.Stderr = &stderr

    //verbose0(fmt.Sprintf("  ----- launch: %s, %v\n", exepath, args))
    //fmt.Println(exepath, args)
    if err = cmd.Start(); err != nil {
      log.Panic(err)                          /* probably shouldn't exit */
    }

    //verbose0(fmt.Sprintf("  ----- launch2: %s, %v\n", exepath, args))
    if err = cmd.Wait(); err == nil {
      res = strings.TrimSpace(stdout.String())

    } else {
      // See if the cmd exited with code != 0
      if exitErr, isExitError := err.(*exec.ExitError); isExitError {
        if status, isWaitStatus := exitErr.Sys().(syscall.WaitStatus); isWaitStatus {
          vvvverbose(fmt.Sprintf("  -- ExitFor(%s): %d\n", cmd.Dir, status.ExitStatus()))

          res = deffault
        }
      }
      //log.Panic(err)                          /* probably shouldn't exit */
    }

    //verbose0(fmt.Sprintf("  ----- launch3: %s, %v\n        %s\n", exepath, args, res))
    out <- res
  }()

  return out, err
}



