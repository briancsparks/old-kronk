package cmd

import (
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "sync"
)

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

var projectDirNames []string

func init() {
  projectDirNames = []string{"AndroidStudioProjects", "EclipseProjects", "GolandProjects", "JordanProjects","vc2019projects","CLionProjects",
    "GameProjects","IdeaProjects","PycharmProjects", "VcProjects", "WebStormProjects"}
}

func repoLocWalk(codeRootsIn []string, stopper func (dirname string, dirs, files []string) ([]string, []string, []string)) (chan entryInfo, chan entryInfo, error) {
  filesChan := make(chan entryInfo)
  dirsChan := make(chan entryInfo)

  codeRoots := keyMirror(codeRootsIn)

  go func() {
    defer close(filesChan)
    defer close(dirsChan)

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

            Verbose(fmt.Sprintf("Processing: %s\n", onePath))

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
            found, moreOf, moreFiles := stopper(onePath, dirs, files)


            // Send files to the channel
            for _, file := range moreFiles {
              filesChan <- entryInfo{name: file, root: root}
            }

            // Send the found items out
            for _, dir := range found {
              dirsChan <- entryInfo{name: dir, root: root}
            }

            // Recurse into the requested sub-dirs
            for _, dir := range moreOf {
              Vverbose(fmt.Sprintf("  -----  MoreOf: %v\n", dir))
              wgOnePath.Add(1)
              go doOnePath(dir, root)
            }

            if len(moreOf) == 0 {
              Vvverbose(fmt.Sprintf("  -----  MoreOf: %v\n", moreOf))
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

// existingCodeDirs builds a slice of strings, where each entry is the 'root' of
// a source tree. Filters out potential dirs that do not exist.
func existingCodeDirs(homeDirs []string) ([]string, error) {
  var result []string

  // Add 'dev', 'projects', and 'go' to each homeDir
  var dirs1 []string
  for _, homeDir := range homeDirs {
    dirs1 = append(dirs1, homeDir)
    dirs1 = append(dirs1, filepath.Join(homeDir, "dev"))
    dirs1 = append(dirs1, filepath.Join(homeDir, "projects"))
    dirs1 = append(dirs1, filepath.Join(homeDir, "go"))
  }

  var dirs2 []string
  for _, dir1 := range dirs1 {
    for _, name := range projectDirNames {
      dirs2 = append(dirs2, filepath.Join(dir1, name))
    }
  }

  // All the dirs we have found
  roots := append(dirs2, dirs1...)

  // Keep the dirs that actually exist
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

