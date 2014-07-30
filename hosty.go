package main

import (
	"fmt"
	"flag"
	"strings"
    "io/ioutil"
    "os"
)

const (
	prefix string = "#hosty-"
	hostsFile string = "/etc/hosts"
	comment string = "#"
)

func main() {
    fileBytes, err := ioutil.ReadFile(hostsFile)
    if err != nil {
		panic(err)
		os.Exit(1)
	}

	fileContent := string(fileBytes)

	entries := make(map[string]string)

    lines := strings.Split(fileContent, "\n")
	for index, line := range lines {
		if strings.HasPrefix(line, prefix) {
			entry := strings.Replace(line, prefix, "", -1)
			nextLineIndex := index + 1
			entries[entry] = lines[nextLineIndex]
		}
	}

	flag.Parse()

	cmd := flag.Arg(0)

	if cmd == "" {
		list(entries)
		return
	}

	switch cmd {
		case "cat":
			fmt.Println(fileContent)
		case "enable":
			entry := flag.Arg(1)
			value := entries[entry]
			if strings.HasPrefix(value, comment) {
				line := strings.Replace(value, comment, " ", 1)
				entries[entry] = line
				fileContent = strings.Replace(fileContent, value, line, 1)
				write(fileContent)
			}
			list(entries)
		case "disable":
			entry := flag.Arg(1)
			value := entries[entry]
			if strings.HasPrefix(value, " ") {
				line := strings.Replace(value, " ", comment, 1)
				entries[entry] = line
				fileContent = strings.Replace(fileContent, value, line, 1)
				write(fileContent)
			}
			list(entries)
	}

	os.Exit(0)
}

func list(entries map[string]string) {
	if len(entries) > 0 {
		fmt.Println("hosty entries:\n")
		index := 1
		for k, v := range entries {
			fmt.Printf("%d) %s\t%s\n", index, k, v)
			index++
		}
	} else {
		fmt.Println("hosty has no entries!")
	}
}

func write(fileContent string) {
	var err = ioutil.WriteFile(hostsFile, []byte(fileContent), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
