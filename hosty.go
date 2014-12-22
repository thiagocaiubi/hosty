package main

import (
	"fmt"
	"flag"
	"strings"
    "io/ioutil"
    "os"
    "sort"
)

const (
	prefix string = "#hosty-"
	hostsFile string = "/etc/hosts"
	comment string = "#"
	whitespace string = " "
	lineBreak  string = "\n"
)

func main() {
	fileContent, entries := read()

	flag.Parse()

	cmd := flag.Arg(0)

	if cmd == "" {
		list(entries)
		os.Exit(0)
	}

	switch cmd {
	case "cat", "c":
		fmt.Println(fileContent)
	case "save", "s":
		if len(flag.Args()) < 4 {
			fmt.Println("hosty bad arguments") //TODO help message
			os.Exit(1)
		}
		entry := flag.Arg(1)
		ip := flag.Arg(2)
		domains := strings.Trim(strings.Join(flag.Args()[3:], whitespace), whitespace)
		newLine := ip + whitespace + domains
		if line, hasEntry := entries[entry]; hasEntry {
			// replacing an existing line will enable it by default
			newLine = whitespace + newLine

			fileContent = strings.Replace(fileContent, line, newLine, 1)
		} else {
			// new entry will be enabled by default
			newLine = whitespace + newLine

			fileContent += prefix + entry + lineBreak
			fileContent += newLine + lineBreak
		}

        write(fileContent)

        entries[entry] = newLine

        list(entries)
	case "enable", "e":
		entry := flag.Arg(1)
		toggle(fileContent, entries, entry, comment, whitespace)
	case "disable", "d":
		entry := flag.Arg(1)
		toggle(fileContent, entries, entry, whitespace, comment)
	}

	os.Exit(0)
}

func list(entries map[string]string) {
	if len(entries) > 0 {
		fmt.Println("hosty entries:\n")
		var keys []string
		for k := range entries {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		index := 0
		for _, k := range keys {
			fmt.Printf("%d) %s\t%s\n", index, k, entries[k])
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

func toggle(fileContent string, entries map[string]string, entry string, current string, replacer string) {
	line := entries[entry]
	if strings.HasPrefix(line, current) {
		newLine := strings.Replace(line, current, replacer, 1)
		entries[entry] = newLine
		fileContent = strings.Replace(fileContent, line, newLine, 1)
		write(fileContent)
	}
	list(entries)
}

func read() (string, map[string]string) {
	fileBytes, err := ioutil.ReadFile(hostsFile)
    if err != nil {
		panic(err)
		os.Exit(1)
	}

	fileContent := string(fileBytes)

	entries := make(map[string]string)

	lines := strings.Split(fileContent, lineBreak)
	for index, line := range lines {
		if strings.HasPrefix(line, prefix) {
			entry := strings.Replace(line, prefix, "", -1)
			nextLineIndex := index + 1
			entries[entry] = lines[nextLineIndex]
		}
	}

	return fileContent, entries
}