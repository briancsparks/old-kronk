package cmd

import "log"

func Verbose0(s string) {
  log.Printf("%s", s)
}

func Verbose(s string) {
  if IsVerbose {
    log.Printf("%s", s)
  }
}

func Vverbose(s string) {
  if IsVverbose {
    log.Printf("%s", s)
  }
}

func Vvverbose(s string) {
  if IsVvverbose {
    log.Printf("%s", s)
  }
}

func Vvvverbose(s string) {
  if IsVvvverbose {
    log.Printf("%s", s)
  }
}

func Verprint(s string) {
  log.Printf("%s", s)
}


