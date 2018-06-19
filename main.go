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
			fmt.Println(markUp(path, regex))
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
			fmt.Println(markUp(exact[i], regex))
			i++
		}
	}
}

func markUp(s string, r *regexp.Regexp) string {
	match := r.FindAllString(filepath.Base(s), -1)
	nomatch := r.Split(filepath.Base(s), -1)
	result := filepath.Dir(s) + "/"
	var in int
	var im int
	lastwasmatched := true
	if r.FindIndex([]byte(filepath.Base(s)))[0] == 0 {
		lastwasmatched = false
	}
	for len(match)+len(nomatch) > im+in {
		if lastwasmatched && len(nomatch) > in {
			result = result + nomatch[in]
			in++
		}
		if !lastwasmatched && len(match) > im {
			result = result + "\033[31;1;4m" + match[im] + "\033[0m"
			im++
		}
		lastwasmatched = !lastwasmatched
	}
	return result
}
