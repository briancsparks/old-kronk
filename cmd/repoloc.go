package cmd

import (
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "sync"
)

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */



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

            verbose(fmt.Sprintf("Processing: %s\n", onePath))

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

            // Call the callback
            found, moreOf := stopper(onePath, dirs, files)


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
              vverbose(fmt.Sprintf("  -----  MoreOf: %v\n", dir))
              go doOnePath(dir, root)
            }

            if len(moreOf) == 0 {
              vvverbose(fmt.Sprintf("  -----  MoreOf: %v\n", moreOf))
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

