package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings package in golang

	// length of string
	fmt.Println(len("Hello World")) // 11

	// Compare
	// Compare returns an integer comparing two strings lexicographically. The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
	fmt.Println(strings.Compare("Hey", "Aey")) // 1
	fmt.Println(strings.Compare("Hey", "Hey")) // 0
	fmt.Println(strings.Compare("Hey", "Zey")) // -1

	// Contains
	// Contains reports whether substr is within s.
	fmt.Println(strings.Contains("Hello World", "World")) // true

	// Count
	// Count counts the number of non-overlapping instances of substr in s. If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
	fmt.Println(strings.Count("deeper", "e")) // 3
	fmt.Println(strings.Count("deeper", "")) // 7

	// EqualFold
	// EqualFold reports whether s and t, interpreted as UTF-8 strings, are equal under Unicode case-folding, which is a more general form of case-insensitivity.
	fmt.Println(strings.EqualFold("AkuL", "akul")) // true

	// Fields
	// Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.
	fmt.Println(strings.Fields("Python Javascript  Rust    Golang   Typescript")) // [Python Javascript Rust Golang Typescript]

	// HasPrefix
	// HasPrefix tests whether the string s begins with prefix.
	fmt.Println(strings.HasPrefix("ProDeveloper", "Pro")) // true

	// HasSuffix
	// HasSuffix tests whether the string s ends with suffix.
	fmt.Println(strings.HasSuffix("String", "ring")) // true

	// Index
	// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
	fmt.Println(strings.Index("this", "i")) // 2

	// IndexAny
	// IndexAny returns the index of the first instance of any Unicode code point from chars in s, or -1 if no Unicode code point from chars is present in s.
	fmt.Println(strings.IndexAny("golang", "a")) // 3

	// Join
	// Join concatenates the elements of its first argument to create a single string. The separator string sep is placed between elements in the resulting string.
	strArr := []string{"Hello", "World", "Everyone"}
	fmt.Println(strings.Join(strArr, " ")) // Hello World Everyone

	// LastIndex
	// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
	fmt.Println(strings.LastIndex("go corona go", "go")) // 10

	// LastIndexAny
	// LastIndexAny returns the index of the last instance of any Unicode code point from chars in s, or -1 if no Unicode code point from chars is present in s.
	fmt.Println(strings.LastIndexAny("golang", "g")) // 5

	// Repeat
	// Repeat returns a new string consisting of count copies of the string s.
	fmt.Println(strings.Repeat("ha", 5)) // hahahahaha

	// Replace
	// Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new. If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string. If n < 0, there is no limit on the number of replacements.
	fmt.Println(strings.Replace("regional emotional grammatical", "al", "ally", 2)) // regionally emotionally grammatical

	// ReplaceAll
	// ReplaceAll returns a copy of the string s with all non-overlapping instances of old replaced by new. If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string.
	fmt.Println(strings.ReplaceAll("regional emotional grammatical", "al", "ally")) // regionally emotionally grammatically

	// Split
	// Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.
	fmt.Println(strings.Split("split with space", " ")) // [split with space]

	// Title
	// Title returns a copy of the string s with all Unicode letters that begin words mapped to their Unicode title case.
	fmt.Println(strings.Title("akul srivastava")) // Akul Srivastava

	// ToLower
	// ToLower returns s with all Unicode letters mapped to their lower case.
	fmt.Println(strings.ToLower("SOME TEXT")) // some text

	// ToUpper
	// ToUpper returns s with all Unicode letters mapped to their upper case.
	fmt.Println(strings.ToUpper("some text")) // SOME TEXT

	// Trim
	// Trim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.
	fmt.Println(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))

}
