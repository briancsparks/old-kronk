package cmd

import "log"

func verbose0(s string) {
  log.Printf("%s", s)
}

func verbose(s string) {
  if Verbose {
    log.Printf("%s", s)
  }
}

func vverbose(s string) {
  if VVerbose {
    log.Printf("%s", s)
  }
}

func vvverbose(s string) {
  if VVVerbose {
    log.Printf("%s", s)
  }
}

func vvvverbose(s string) {
  if VVVVerbose {
    log.Printf("%s", s)
  }
}

func verprint(s string) {
  log.Printf("%s", s)
}


