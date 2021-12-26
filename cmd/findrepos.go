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
)

// findreposCmd represents the findrepos (kronk code findrepos) command
var findreposCmd = &cobra.Command{
	Use:   "findrepos",
	Short: "Find and list all repos",
	Long: `Find and list all repos

Looks in all the usual places and lists what repo is where.`,

	Run: func(cmd *cobra.Command, args []string) {
    findreposInitForSub()
    //log.Printf("findr1: %v, %v, %v\n", IsVerbose, IsVverbose, IsVvverbose)

    Verbose(fmt.Sprintln("findrepos called"))
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
  super  *superDir
  name    string
  root    string
}

func find() {
  userRepoRoots := os_UserHomeDirs()
  userRepoRoots = append(userRepoRoots, "D:\\data")

  if IsVerbose {
    for _, roots := range userRepoRoots {
      Verprint(fmt.Sprintf("Root: %s\n", roots))
    }
  }

  codeRoots, err := existingCodeDirs(userRepoRoots)
  Check(err)

  Vverbose(fmt.Sprintf("codeRoots: %v\n", codeRoots))

  shouldStop := func (dirname string, shortDirsIn, shortFilesIn []string) ( /*found*/ []string, /*moreOf*/ []string, /*moreFiles*/ []string) {

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

  // Start the walk
  files, dirs, err := repoLocWalk(codeRoots, shouldStop)
  Check(err)

  i := 10
  var wg sync.WaitGroup

  // Consume all files found
  wg.Add(1)
  go func() {
    defer wg.Done()
    for ent := range files {
      i *= 1
      _ = ent
    }
  }()

  // Consume dirs found
  wg.Add(1)
  go func() {
    defer wg.Done()

    for dir := range dirs {
      wg.Add(1)
      go func(dir entryInfo) {
        defer wg.Done()

        // Do something with the dir
        checkDir(dir)
      }(dir)
    }
  }()

  wg.Wait()
}

func checkDir(dir entryInfo) {
  Verbose(fmt.Sprintf("  ---- checkDir: %s\n", dir.name))
  gitConfigFile := filepath.Join(dir.name, ".git", "config")

  _, err := os.Stat(gitConfigFile)
  if err == nil || os.IsExist(err) {

    gitArgs := []string{"config", "--get", "remote.origin.url"}

    output, err := launchForResult("git", gitArgs, dir.name, "")
    Check(err)

    originUrl := <-output
    if len(originUrl) > 0 {
      if strings.Contains(originUrl, "briancsparks") {

        gitStatusChan, err := launchForResult("git", []string{"status", "--short"}, dir.name, "")
        Check(err)

        gitStatus := <-gitStatusChan

        if len(gitStatus) > 0 {
          fmt.Printf("%-94s %s\n", originUrl, cygpath(dir.name))
          fmt.Printf("=======================\n%s\ngit status: %s\n------------------------\n", dir.name, gitStatus)
        }
      }
    }
  }
}


