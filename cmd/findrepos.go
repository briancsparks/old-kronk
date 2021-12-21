package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

// TODO: Fix ending soon.
// TODO: Look into .git/config to know that it's _my_ repo.
// TODO: Show the github url / sort by github url

import (
  "fmt"
  "github.com/spf13/cobra"
  "path/filepath"
  "strings"
)

// findreposCmd represents the findcode command
var findreposCmd = &cobra.Command{
	Use:   "findrepos",
	Short: "Find and list all repos",
	Long: `Find and list all repos

Looks in all the usual places and lists what repo is where.`,

	Run: func(cmd *cobra.Command, args []string) {
    findreposInitForSub()
    //log.Printf("findr1: %v, %v, %v\n", Verbose, VVerbose, VVVerbose)

    verbose(fmt.Sprintln("findrepos called"))
    find()
	},
}

func init() {
	codeCmd.AddCommand(findreposCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findreposCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}

func findreposInitForSub() {
  codeInitForSub()
}


var lessof map[string]string

func init() {
  lessof = keyMirror(strings.Split("node_modules,.git", ","))
}

type entryInfo struct {
  name string
  root string
}

func find() {
  userRepoRoots := os_UserHomeDirs()
  userRepoRoots = append(userRepoRoots, "D:\\data")

  if Verbose {
    for _, roots := range userRepoRoots {
      verprint(fmt.Sprintf("Root: %s\n", roots))
    }
  }

  codeRoots, err := codeDirs(userRepoRoots)
  check(err)

  vverbose(fmt.Sprintf("codeRoots: %v\n", codeRoots))

  shouldStop := func (dirname string, shortDirsIn, shortFilesIn []string) ( /*found*/ []string, /*moreOf*/ []string) {
    dirsIn := smap(shortDirsIn, prependPath(dirname))
    filesIn := smap(shortFilesIn, prependPath(dirname))
    dirs := keyMirror(dirsIn)
    files := keyMirror(filesIn)

    if _, ok := dirs[filepath.Join(dirname, ".git")]; ok {
      return slc(dirname), []string{}
    }

    if _, ok := files[filepath.Join(dirname, "package.json")]; ok {
      return slc(dirname), []string{}
    }

    moreOf := make([]string, 0)
    for _, shortDir := range shortDirsIn {
      if _, ok := lessof[shortDir]; !ok {
        moreOf = append(moreOf, filepath.Join(dirname, shortDir))
      }
    }

    return []string{}, moreOf
  }

  files, dirs, err := superWalk(codeRoots, shouldStop)
  check(err)

  for i := 0;; i++ {
    select {
    case dir := <- dirs:
     i *= 1
     fmt.Printf("%s\n", dir.name)
    case <-files:
      i *= 1
      //fmt.Printf("File: %s\n", file.name)
    }
  }

}


