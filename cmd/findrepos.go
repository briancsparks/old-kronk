package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

// TODO: Fix ending soon.
// TODO: Look into .git/config to know that it's _my_ repo.
// TODO: Show the github url / sort by github url

import (
  "fmt"
  "github.com/spf13/cobra"
  "io/ioutil"
  "os"
  "path/filepath"
  "sync"
)

// findcodeCmd represents the findcode command
var findcodeCmd = &cobra.Command{
	Use:   "findrepos",
	Short: "Find and list all repos",
	Long: `Find and list all repos

Looks in all the usual places and lists what repo is where.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("findrepos called")
    find()
	},
}

func init() {
	codeCmd.AddCommand(findcodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findcodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findcodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type entryInfo struct {
  name string
  root string
}

func find() {
  userRepoRoots := os_UserHomeDirs()
  userRepoRoots = append(userRepoRoots, "D:\\data")

  for _, roots := range userRepoRoots {
    fmt.Println(roots)
  }

  codeRoots, err := codeDirs(userRepoRoots)
  check(err)

  shouldStop := func (dirname string, shortDirsIn, shortFilesIn []string) ( /*found*/ []string, /*moreOf*/ []string) {
    dirsIn := smap(shortDirsIn, func(s string) string { return filepath.Join(dirname, s) })
    filesIn := smap(shortFilesIn, func(s string) string { return filepath.Join(dirname, s) })
    dirs := keyMirror(dirsIn)
    files := keyMirror(filesIn)
    //fmt.Printf("stopper: %s\n    dirsIn:   %v\n    filesIn: %v\n    %v\n    %v\n", q(dirname), shortDirsIn, shortFilesIn, dirs, qq(dirsIn))

    wantMore := true
    found := make(map[string]string)

    if _, ok := dirs[filepath.Join(dirname, ".git")]; ok {
      found[dirname] = dirname
      wantMore = false
    }

    if _, ok := files[filepath.Join(dirname, "package.json")]; ok {
      found[dirname] = dirname
      wantMore = false
    }

    if !wantMore {
      return keys(found), []string{}
    }

    moreOf := make([]string, 0)
    for _, shortDir := range shortDirsIn {
      if shortDir == "node_modules" {
        continue
      }

      dir := filepath.Join(dirname, shortDir)

      //if len(dir) <= 55 {
        moreOf = append(moreOf, dir)
      //} else {
      // fmt.Printf("stop %s\n", q(dir))
      //}
    }

    return keys(found), moreOf
  }

  files, dirs, err := superWalk(codeRoots, shouldStop)
  check(err)

  for i := 0;; i++ {
    select {
    case dir := <- dirs:
     i *= 1
     fmt.Printf("Dir:  %s\n", dir.name)
    case <-files:
      i *= 1
      //fmt.Printf("File: %s\n", file.name)
    }
  }

}

func superWalk(codeRootsIn []string, stopper func (dirname string, dirs, files []string) ([]string, []string)) (chan entryInfo, chan entryInfo, error) {
  filesChan := make(chan entryInfo)
  dirsChan := make(chan entryInfo)

  codeRoots := keyMirror(codeRootsIn)

  go func() {
    defer close(filesChan)

    var wgCodeRoots sync.WaitGroup
    for _, root := range codeRootsIn {
      func() {
        doOnePath := func(onePath, rootDir string) {}

        wgCodeRoots.Add(1)
        go func(root string) {
          defer wgCodeRoots.Done()

          // ------------------
          var wgOnePath sync.WaitGroup
          doOnePath = func(onePath, rootDir string) {
            defer wgOnePath.Done()

            // Make sure we only do it once
            if onePath != rootDir {
              if _, ok := codeRoots[onePath]; ok {
                //fmt.Printf("Skip: %s (%s)\n", onePath, rootDir)
                return
              }
            }

            files := make([]string, 0)
            dirs := make([]string, 0)
            entries, _ := ioutil.ReadDir(onePath)
            for _, entry := range entries {
              if ! entry.IsDir() {
                files = append(files, entry.Name())
              } else {
                dirs = append(dirs, entry.Name())
              }
            }

            found, moreOf := stopper(onePath, dirs, files)
            //fmt.Printf("stopped: %s\n    found: %v\n    moreOf: %v\n", q(onePath), qq(found), qq(moreOf))

            // Send files to the channel
            for _, file := range files {
              filesChan <- entryInfo{name: file, root: root}
            }

            for _, dir := range found {
              dirsChan <- entryInfo{name: dir, root: root}
            }

            // Send the requested sub-dirs
            for _, dir := range moreOf {
              wgOnePath.Add(1)
              go doOnePath(dir, root)
            }
          }
          // ------------------

          wgOnePath.Add(1)
          go doOnePath(root, root)

          wgOnePath.Wait()
        }(root)
      }()
    }

    wgCodeRoots.Wait()
  }()

  return filesChan, dirsChan, nil
}

func codeDirs(homeDirs []string) ([]string, error) {
  var result []string

  var dirs1 []string
  for _, homeDir := range homeDirs {
    dirs1 = append(dirs1, homeDir)
    dirs1 = append(dirs1, filepath.Join(homeDir, "dev"))
    dirs1 = append(dirs1, filepath.Join(homeDir, "projects"))
    dirs1 = append(dirs1, filepath.Join(homeDir, "go"))
  }

  var projectDirNames = []string{"AndroidStudioProjects", "EclipseProjects", "GolandProjects", "JordanProjects","vc2019projects","CLionProjects",
                                 "GameProjects","IdeaProjects","PycharmProjects", "VcProjects", "WebStormProjects"}

  var dirs2 []string
  for _, dir1 := range dirs1 {
    for _, name := range projectDirNames {
      dirs2 = append(dirs2, filepath.Join(dir1, name))
    }
  }

  roots := append(dirs2, dirs1...)

  dirs3 := dirs2[:0]
  for _, root := range roots {
    _, err := os.Stat(root)
    if err == nil || os.IsExist(err) {
      dirs3 = append(dirs3, root)
    }
  }

  result = dirs3
  return result, nil
}




