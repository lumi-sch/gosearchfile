package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:")
		fmt.Println("	searchfile <filename>")
		os.Exit(1)
	}
	var exact []string
	regex, err := regexp.Compile(os.Args[1])
	if err != nil {
		fmt.Println("Invalid regular expression.")
		os.Exit(1)
	}
	_ = filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
		}
		if info.Name() == "proc" {
			return filepath.SkipDir
		}
		if regex.MatchString(info.Name()) {
			fmt.Println(path)
		}
		if os.Args[1] == info.Name() {
			exact = append(exact, path)
		}
		return nil
	})
	if len(exact) > 0 {
		fmt.Println("-----------------------")
		fmt.Println("|    EXACT MATCHES    |")
		fmt.Println("-----------------------")
		i := 0
		for len(exact) > i {
			fmt.Println(exact[i])
			i++
		}
	}
}
