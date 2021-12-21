package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

// TODO: Fix ending soon.
// TODO: Look into .git/config to know that it's _my_ repo.
// TODO: Show the github url / sort by github url

import (
  "fmt"
  "github.com/spf13/cobra"
  "os"
  "path/filepath"
  "strings"
  "sync"
  "time"
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

  t :=time.Now()

  shouldStop := func (dirname string, shortDirsIn, shortFilesIn []string) ( /*found*/ []string, /*moreOf*/ []string, /*moreFiles*/ []string) {
    if time.Since(t).Seconds() > 1 {
      vvvverbose(fmt.Sprintf("  ----- Looking at %s\n", dirname))
      t = time.Now()
    }

    dirsIn := smap(shortDirsIn, prependPath(dirname))
    filesIn := smap(shortFilesIn, prependPath(dirname))
    dirs := keyMirror(dirsIn)
    files := keyMirror(filesIn)

    if _, ok := dirs[filepath.Join(dirname, ".git")]; ok {
      return slc(dirname), []string{}, []string{}
    }

    if _, ok := files[filepath.Join(dirname, "package.json")]; ok {
      return slc(dirname), []string{}, []string{}
    }

    moreOf := make([]string, 0)
    for _, shortDir := range shortDirsIn {
      if _, ok := lessof[shortDir]; !ok {
        moreOf = append(moreOf, filepath.Join(dirname, shortDir))
      }
    }

    return []string{}, moreOf, []string{}
  }

  files, dirs, err := superWalk(codeRoots, shouldStop)
  check(err)

  i := 10
  var wg sync.WaitGroup

  wg.Add(1)
  go func() {
    defer wg.Done()
    for ent := range files {
      i *= 1
      _ = ent
    }
  }()

  wg.Add(1)
  go func() {
    defer wg.Done()

    for dir := range dirs {
      wg.Add(1)
      go func(dir entryInfo) {
        defer wg.Done()
        checkDir(dir)
      }(dir)
    }
  }()


  //for i := 0;; i++ {
  //  select {
  //  case dir := <- dirs:
  //   i *= 1
  //   //fmt.Printf("%s\n", dir.name)
  //   wg.Add(1)
  //   go func() {
  //     defer wg.Done()
  //     checkDir(dir)
  //   }()
  //
  //  case <-files:
  //    i *= 1
  //    //fmt.Printf("File: %s\n", file.name)
  //  }
  //}

  wg.Wait()
}

func checkDir(dir entryInfo) {
  verbose(fmt.Sprintf("  ---- checkDir: %s\n", dir.name))
  gitConfigFile := filepath.Join(dir.name, ".git", "config")

  _, err := os.Stat(gitConfigFile)
  if err == nil || os.IsExist(err) {
    //verbose0(fmt.Sprintf("gitConfigFile: %s\n", gitConfigFile))

    gitArgs := []string{"config", "--get", "remote.origin.url"}

    output, err := launchForResult("git", gitArgs, dir.name, "")
    check(err)

    originUrl := <-output
    if len(originUrl) > 0 {
      fmt.Printf("%-104s %s\n", originUrl, dir.name)
      //fmt.Printf("gitUrl: %s\n        %s\n", originUrl, dir.name)
    //} else {
    //  fmt.Printf("Empty:  %s\n", dir.name)
    }
  }
}


