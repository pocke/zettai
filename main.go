package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	paths := make([]string, 0)
	for _, v := range os.Args[1:] {
		p, err := filepath.Abs(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		paths = append(paths, p)
	}
	fmt.Print(strings.Join(paths, " "))
}
