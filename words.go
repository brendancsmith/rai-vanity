package main

import (
	"fmt"
	"os"
	"strings"

	"./itertools"
	//"github.com/derekparker/trie"
)

// CorpusOptions : settings for reading a body of work into a word list
type CorpusOptions struct {
	Charset       string
	IgnoreCase    bool
	SquashHyphens bool
	SquashSpaces  bool
}

// LoadWords : load a corpus file
func LoadWords(resourcePath string, opts *CorpusOptions) []string {
	lines := readLines(resourcePath)
	lines = cleanList(lines, opts)
	return lines
}

func readLines(resourcePath string) []string {
	corpusBytes, err := Asset(resourcePath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lines := strings.FieldsFunc(
		string(corpusBytes),
		func(char rune) bool {
			return (char == '\n' || char == '\r')
		})

	return lines
}

func cleanList(lines []string, opts *CorpusOptions) []string {
	// trim whitespace
	lines = itertools.Map(lines, func(s string) string {
		return strings.TrimSpace(s)
	})

	// remove empty lines
	lines = itertools.Filter(lines, func(s string) bool {
		return len(s) != 0
	})

	// remove spaces and hyphens
	lines = itertools.Map(lines, func(s string) string {
		s = strings.Replace(s, " ", "", -1)
		s = strings.Replace(s, "-", "", -1)
		return s
	})

	// lower case
	lines = itertools.Map(lines, func(s string) string {
		return strings.ToLower(s)
	})

	// filter against charset
	lines = itertools.Filter(lines, func(s string) bool {
		for _, c := range s {
			if !strings.ContainsRune(opts.Charset, c) {
				return false
			}
		}
		return true
	})

	return lines
}
