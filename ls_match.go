package main

import "fmt"
import "flag"
import "path/filepath"
import "strings"
import "os"
import "regexp"

var options struct {
	mask  string
	first bool
}

func print(path string) {
	fmt.Println(path)

	if options.first { os.Exit(0) }
}

func substituteHomeDir(path string) string {
	if os.Getenv("HOME") != "" {
		if path == "~" {
			return os.Getenv("HOME")
		} else if path[0:2] == "~" + string(os.PathSeparator) {
			return filepath.Join(os.Getenv("HOME"), path[2:])
		}
	}

	return path
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
	options.mask = regexp.MustCompile(`%\d+`).ReplaceAllString(options.mask, "")

	for _, next_mask := range filepath.SplitList(options.mask) {
		paths, _ := filepath.Glob(substituteHomeDir(next_mask))

		for _, path := range paths {
			if stat, err := os.Stat(path); err == nil && stat.IsDir() {
				print(path)
			}
		}
	}
}
