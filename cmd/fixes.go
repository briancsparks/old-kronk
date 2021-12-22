package cmd

import "os"

var (
  fixed_UserHomeDir string
  fixed_UserCacheDir string
  fixed_UserConfigDir string
)

func init() {
  var err error
  fixed_UserHomeDir, err = os.UserHomeDir()
  Check(err)
  fixed_UserCacheDir, err = os.UserCacheDir()
  Check(err)
  fixed_UserConfigDir, err = os.UserConfigDir()
  Check(err)
}

func os_UserHomeDir() string {
  return fixed_UserHomeDir
}

func os_UserHomeDirs() []string {
  return []string{fixed_UserHomeDir}
}

func os_UserCacheDir() string {
  return fixed_UserCacheDir
}

func os_UserConfigDir() string {
  return fixed_UserConfigDir
}

