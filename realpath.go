package main

import (
	"os"
	"fmt"
	"flag"
	"path/filepath"
)

func main() {
	var quiet bool
	flag.BoolVar(&quiet, "q", false, "warnings will not be printed")

	flag.Parse()

	for _, path := range flag.Args() {
		s, err := filepath.Abs(path)
		if err != nil {
			if !quiet {
				fmt.Fprintln(os.Stderr, err)
			}
			continue
		}
		s, err = filepath.EvalSymlinks(s)
		if err != nil {
			if !quiet {
				fmt.Fprintln(os.Stderr, err)
			}
			continue
		}
		fmt.Println(s)
	}
}
