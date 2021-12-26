package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "path/filepath"
  "strings"
  "sync"

  "github.com/spf13/cobra"
)

// gitsstatusCmd represents the gitsstatus command
var gitsstatusCmd = &cobra.Command{
	Use:   "gitsstatus",
  Short: "git status for all repos",
  Long: `git status for all repos

Looks in all the usual repo places and does git status on them.`,

  Run: func(cmd *cobra.Command, args []string) {
    gitsstatusInitForSub()
		fmt.Println("gitsstatus called")

    gitsstatus()
	},
}

func init() {
	codeCmd.AddCommand(gitsstatusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gitsstatusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gitsstatusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func gitsstatusInitForSub() {
  codeInitForSub()
}

func gitsstatus() {
  userRepoRoots := os_UserHomeDirs()
  userRepoRoots = append(userRepoRoots, "D:\\data")
  Verbose(fmt.Sprintf("userRepoRoots: %v\n", userRepoRoots))

  //fmt.Printf("Verbosezz: v: %v, vv: %v, vvv: %v, vvvv: %v\n", IsVerbose, IsVverbose, IsVvverbose, IsVvvverbose)

  codeRoots, err := superExistingCodeDirs(userRepoRoots)
  CheckMsg(err, "superExistingCodeDirs")

  Vverbose(fmt.Sprintf("codeRoots: %v\n", codeRoots))

  stopper := func(super *superDir) (/*foundFiles*/ []string, /*foundDirs*/ []string, /*moreOf*/ []string) {
    //Verbose0(fmt.Sprintf("------------stopper: dir=%s\n", super.fulldirpath))
    //Verbose0(fmt.Sprintf("------------stopper: dir=%s\n -- dirs: %v\n", super.fulldirpath, super.dirs))

    moreOf := make([]string, 0)

    if super.hasDir(".git") {
      //Verbose0(fmt.Sprintf("====================== stopper hasdir .git\n"))
      return []string{}, []string{super.fulldirpath}, moreOf
    }

    if super.hasFile("package.json") {
      //Verbose0(fmt.Sprintf("====================== stopper hasfile package.json\n"))
      return []string{}, []string{super.fulldirpath}, moreOf
    }

    for _, shortDir := range super.dirs {
      if _, ok := lessof[shortDir]; !ok {
        //Verbose0(fmt.Sprintf("====================== stopper(%s) moreof= %s\n", super.fulldirpath, shortDir))
        moreOf = append(moreOf, filepath.Join(super.fulldirpath, shortDir))
      }
    }

    //Verbose0(fmt.Sprintf("====================== stopper normal\n"))
    return []string{}, []string{}, moreOf
  }

  // Start the walk
  files, dirs, err := superWalk(codeRoots, stopper)
  Check(err)

  //Verbose0(fmt.Sprintf("Have chans from superWalk\n"))

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

    func() {
      for dir := range dirs {
        wg.Add(1)
        go func(dir entryInfo) {
          defer wg.Done()

          // Do something with the dir
          gitsStatusDir(&dir)
        }(dir)
      }
    }()
  }()

  Verbose0(fmt.Sprintf("Main waiting on wg\n"))
  wg.Wait()
  Verbose0(fmt.Sprintf("Main done waiting on wg\n"))

}

func gitsStatusDir(entry *entryInfo) {
  dir  := entry.super
  //Verbose0(fmt.Sprintf("  ---- gitsStatusDir: %s\n", dir.fulldirpath))

  originUrlChan, err := dir.launchForResult("git", []string{"config", "--get", "remote.origin.url"}, "")
  Check(err)

  originUrl := <-originUrlChan

  if len(originUrl) > 0 {
   if strings.Contains(originUrl, "briancsparks") {
     //Verbose0(fmt.Sprintf("Checking %s\n", entry.super.fulldirpath))

     gitStatusChan, err := dir.launchForResult("git", []string{"status", "--short"}, "")
     Check(err)

     gitStatus := <-gitStatusChan

     if len(gitStatus) > 0 {
       fmt.Printf("%-94s %s\n", originUrl, cygpath(dir.fulldirpath))
       fmt.Printf("=======================\n%s\ngit status: %s\n------------------------\n", dir.fulldirpath, gitStatus)
     }
   }
  }

}


