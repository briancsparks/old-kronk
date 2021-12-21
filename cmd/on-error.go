package cmd

import (
  "log"
)

const (
  OnErrorPanic = iota
  OnErrorFatal
  OnErrorLogAndContinue
  OnErrorPropigate                  /* Log and return the err */
)

var mode int
var modeOnDeadlyError int

func init() {
  mode = OnErrorPropigate
  modeOnDeadlyError = OnErrorPanic
}

func onError(newMode int) {
  mode = newMode
}

func onDeadly(newMode int) {
  modeOnDeadlyError = newMode
}

func check(err error) {
  if err != nil {
    if modeOnDeadlyError == OnErrorPanic {
      log.Panicln(err)
      return
    }

    if modeOnDeadlyError == OnErrorFatal {
      log.Fatalln(err)
      return
    }

    // -------------
    // Calling deadly check, but deadly mode is not lethal

    // deadly check, but want log-and-continue... for temporarily blasting past situations
    // that are not handled yet, and are not really deadly (like in active development.)
    if modeOnDeadlyError == OnErrorLogAndContinue {
      log.Println(err)
      return
    }

    // Deadly check, but want to propigate the error... not really possible, but the most
    // reasonable thing is panic.
    if modeOnDeadlyError == OnErrorLogAndContinue || modeOnDeadlyError == OnErrorPropigate {
      log.Panicln(err)
      return
    }
  }
}

func check2(err error) bool {
  if err != nil {
    if mode == OnErrorPanic {
      log.Panicln(err)
      return true
    }

    if mode == OnErrorFatal {
      log.Fatalln(err)
      return true
    }

    if mode == OnErrorLogAndContinue {
      log.Println(err)
      return false
    }

    if mode == OnErrorPropigate {
      log.Println(err)
      return true
    }

  }

  return false
}

