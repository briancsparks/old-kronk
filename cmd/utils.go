package cmd

import "log"

func checkU(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

