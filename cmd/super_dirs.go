package cmd

import (
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "strings"
  "sync"
)

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

//var projectDirNames []string
//
//func init() {
//  projectDirNames = []string{"AndroidStudioProjects", "EclipseProjects", "GolandProjects", "JordanProjects","vc2019projects","CLionProjects",
//    "GameProjects","IdeaProjects","PycharmProjects", "VcProjects", "WebStormProjects"}
//}

type superDir struct {
  fulldirpath   string
  dirs          []string        /* short names */
  files         []string        /* short names */

  root          string

  dmap          map[string]bool
  fmap          map[string]bool
}

func newSuperDir(dirpath, root string) *superDir {
  //Verbose0(fmt.Sprintf("  ----newSuperDirIn----\n  -- %s\n  -- root= %s\n", dirpath, root))

  var super superDir
  super.dirs = make([]string, 0)
  super.files = make([]string, 0)
  super.dmap = make(map[string]bool)
  super.fmap = make(map[string]bool)

  super.fulldirpath = dirpath
  super.root        = root

  entries, _ := ioutil.ReadDir(super.fulldirpath)
  //Verbose0(fmt.Sprintf("ioutil.ReadDir dir= %s,    err= %v\n", super.fulldirpath, err))
  //CheckMsg(err, fmt.Sprintf("newSuperDir(%s)", dirpath))

  for _, entry := range entries {
    if entry.IsDir() {
      super.dirs = append(super.dirs, entry.Name())
    } else {
      super.files = append(super.files, entry.Name())
    }
  }

  //Verbose0(fmt.Sprintf("  ----newSuperDir----\n  -- %s\n  -- root= %s\n  -- d: %v\n  -- f: %v\n", super.fulldirpath, super.root, super.dirs, super.files))
  return &super
}

func setify(m map[string]bool, entries []string, fulldirpath string) {
  if len(m) == 0 && len(entries) != 0 {
    for _, entry := range entries {
      m[entry] = true
      m[filepath.Join(fulldirpath, entry)] = true
    }
  }
}

func (super *superDir) hasDir(s string) bool {

  // Ensure ok
  setify(super.dmap, super.dirs, super.fulldirpath)

  if _, ok := super.dmap[s]; ok {
    return true
  }

  return false
}

func (super *superDir) hasFile(s string) bool {

  // Ensure ok
  setify(super.fmap, super.files, super.fulldirpath)

  if _, ok := super.fmap[s]; ok {
    return true
  }

  return false
}

func (super *superDir) launchForResult(exename string, args []string, deffault string)  (chan string, error) {
  return launchForResult(exename, args, super.fulldirpath, deffault)
}

//func repoLocWalk(codeRootsIn []string, stopper func (dirname string, dirs, files []string) ([]string, []string, []string)) (chan entryInfo, chan entryInfo, error) {
func superWalk(topRootsIn []string, stopper func(*superDir) ([]string, []string, []string)) (/*files*/ chan entryInfo, /*dirs*/ chan entryInfo, error) {
  filesChan := make(chan entryInfo)
  dirsChan := make(chan entryInfo)

  topRoots := keyMirror(topRootsIn)

  go func() {
    defer close(filesChan)
    defer close(dirsChan)

    var wgTopRoots sync.WaitGroup
    for _, root := range topRootsIn {
      func() {
        doOnePath := func(onePath, rootDir string) {}

        wgTopRoots.Add(1)
        go func(root string) {
          defer wgTopRoots.Done()

          // ------------------
          var wgOnePath sync.WaitGroup
          doOnePath = func(onePath, rootDir string) {
            defer wgOnePath.Done()

            Verbose(fmt.Sprintf("Processing(super): %s, rootDir: %s, root: %s\n", onePath, rootDir, root))

            // Make sure we only do it once
            if onePath != rootDir {
              if _, ok := topRoots[onePath]; ok {
                //fmt.Printf("Skip: %s (%s)\n", onePath, rootDir)
                return
              }
            }

            //entries, _ := ioutil.ReadDir(onePath)
            super := newSuperDir(onePath, root)

            foundFiles, foundDirs, moreOf := stopper(super)
            Verbose(fmt.Sprintf("Stop res:\n  -- %s\n  -- foundFiles: %v\n  -- foundDirs: %v\n  -- moreOf: %v\n", onePath, foundFiles, foundDirs, moreOf))

            for _, file := range foundFiles {
              filesChan <- entryInfo{super:super, name:file, root:root}
            }

            for _, dir := range foundDirs {
              dirsChan <- entryInfo{super:super, name:dir, root:root}
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

          //Verbose0(fmt.Sprintf("onePath Waiting, root: %s\n", root))
          wgOnePath.Wait()

          //Verbose0(fmt.Sprintf("onePath Done, root: %s\n", root))
        }(root)
      }()
    }

    //Verbose0(fmt.Sprintf("superWalk Waiting\n"))
    wgTopRoots.Wait()
    //Verbose0(fmt.Sprintf("superWalk Done Waiting\n"))
  }()

  return filesChan, dirsChan, nil
}

func superExistingCodeDirs(dirs []string) ([]string, error) {
 roots := make([]string, 0)

 roots = superSubDirs(dirs, strings.Split("dev,projects,go", ","))
 roots = superSubDirs(roots, projectDirNames)

 result := make([]string, 0)
 for _, root := range roots {
   _, err := os.Stat(root)

   // TODO: return error on real error
   if err == nil || os.IsExist(err) {
     result = append(result, root)
   }
 }

 return result, nil
}

// superSubDirs appends each of subs to the end of each of dirs, using filepath.Join()
func superSubDirs(dirs []string, subs []string) []string {
  result := make([]string, 0)

  for _, dir := range dirs {
    result = append(result, dir)
    for _, sub := range subs {
      result = append(result, filepath.Join(dir, sub))
    }
  }

  Vvverbose(fmt.Sprintf("sSubDirs: %v, subs: %v\n\n   --- %v", dirs, subs, result))

  return result
}


