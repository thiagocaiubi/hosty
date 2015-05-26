package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

const (
	prefix     string = "#hosty-"
	hostsFile  string = "/etc/hosts"
	comment    string = "#"
	whitespace string = " "
	lineBreak  string = "\n"
	empty      string = ""
)

func main() {
	fileContent := read()
	entries := parseEntries(fileContent)

	flag.Parse()

	cmd := flag.Arg(0)

	if cmd == empty {
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

		newLine := save(fileContent, entries, entry, ip, domains)

		entries[entry] = newLine

		list(entries)
	case "enable", "e":
		entry := flag.Arg(1)
		toggle(fileContent, entries, entry, comment, whitespace)
	case "disable", "d":
		entry := flag.Arg(1)
		toggle(fileContent, entries, entry, whitespace, comment)
	case "remove", "r":
		entry := flag.Arg(1)
		if line, hasEntry := entries[entry]; hasEntry {
			fileContent = strings.Replace(fileContent, prefix+entry+lineBreak, empty, 1)
			fileContent = strings.Replace(fileContent, line+lineBreak, empty, 1)

			write(fileContent)

			delete(entries, entry)

			list(entries)
		} else {
			fmt.Println("hosty has no entry: " + entry)
			os.Exit(1)
		}
	}

	os.Exit(0)
}

// list prints pretty entries output
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

// write fileContent to hostsFile
func write(fileContent string) {
	var err = ioutil.WriteFile(hostsFile, []byte(fileContent), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//TODO toggle should be self contained about how char to replace
// toggle change entry's status from enabled to disabled and the other way around
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

// read hostsFile
// return fileContent as string
func read() string {
	fileBytes, err := ioutil.ReadFile(hostsFile)
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	return string(fileBytes)
}

// parse hostsFile content and put managed entries in a map
// return entries' map[string]string
func parseEntries(fileContent string) map[string]string {
	entries := make(map[string]string)

	lines := strings.Split(fileContent, lineBreak)
	for index, line := range lines {
		if strings.HasPrefix(line, prefix) {
			entry := strings.Replace(line, prefix, empty, -1)
			nextLineIndex := index + 1
			entries[entry] = lines[nextLineIndex]
		}
	}

	return entries
}

// update new entry if it already exists
// save new entry to fileContent
// invoke write
func save(fileContent string, entries map[string]string, entry string, ip string, domains string) string {
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

	return newLine
}
