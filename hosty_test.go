package main

import (
	"reflect"
	"testing"
)

func TestParseEntriesReturnsEntries(t *testing.T) {
	fileContent := "#hosty-example\n127.0.0.1 example.com"
	entries := parseEntries(fileContent)

	entriesExpected := map[string]string{
		"example": "127.0.0.1 example.com",
	}

	if !reflect.DeepEqual(entriesExpected, entries) {
		t.Errorf("Unexpected entries: %v", entries)
	}
}
