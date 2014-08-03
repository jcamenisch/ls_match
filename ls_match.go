package main

import "fmt"
import "flag"
import "path/filepath"
import "strings"
import "os"

var options struct {
  mask  string
  first bool
}

func print(path string) {
  fmt.Println(path)

  if options.first { os.Exit(0) }
}

func main() {
  flag.StringVar(&options.mask, "mask",  "",    "A custom glob mask for looking up your directory")
  flag.BoolVar( &options.first, "first", false, "Only print the first match found")
  flag.Parse()

  for i, arg := range flag.Args() {
    token := fmt.Sprintf("%%%d", i + 1)

    if strings.Contains(options.mask, token) {
      options.mask = strings.Replace(options.mask, token, arg, -1)
    }
  }

  for _, next_mask := range filepath.SplitList(options.mask) {
    matches, _ := filepath.Glob(next_mask)

    for _, match := range matches {
      print(match)
    }
  }
}
