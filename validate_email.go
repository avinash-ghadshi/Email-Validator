package main

import (
    "os"
    "fmt"
    "bufio"
    "regexp"
    "strings"
    "runtime"
    _"io/ioutil"
    "path/filepath"
)

func main() {
	_, file, _, _ := runtime.Caller(0)
	if len(os.Args) != 2 {
		fmt.Println("USAGE: go run "+file+" <filename>")
		os.Exit(1)
	}

	if _, err := os.Stat(os.Args[1]); err != nil {
		fmt.Println(os.Args[1]+" does not exist")
		os.Exit(1)
	}

	FILENAME := os.Args[1]
	//fmt.Println(FILENAME)
	f, err := os.Open(FILENAME)
	if err != nil {
		fmt.Println("File reading error", err)
		os.Exit(1)
	}
	defer f.Close()

	correct := make([]string, 0, 32*1024)
	incorrect := make([]string, 0, 32*1024)
	var validID = regexp.MustCompile(`^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		if validID.MatchString(scanner.Text()) == true {
			correct = append(correct, scanner.Text())
			continue
		}
		incorrect = append(incorrect, scanner.Text())

	}
	ct := strings.Join(correct, "\n")
	ict := strings.Join(incorrect, "\n")

	dir := filepath.Dir(FILENAME)

	cf := dir+"/correct_emails.txt"
	icf := dir+"/incorrect_emails.txt"

	cfile, err := os.Create(cf)
	if err != nil {
		panic(err)
	}
	icfile, err := os.Create(icf)
	if err != nil {
		panic(err)
	}
	cwriter := bufio.NewWriter(cfile)
	icwriter := bufio.NewWriter(icfile)
	_, err = cwriter.WriteString(ct)
	_, err = icwriter.WriteString(ict)
	cwriter.Flush()
	icwriter.Flush()

	fmt.Printf("Correct Email IDs File: %s\tCount:%d\n", cf, len(correct))
	fmt.Printf("Incorrect Email IDs File: %s\tCount:%d\n", icf, len(incorrect))
	return
}
