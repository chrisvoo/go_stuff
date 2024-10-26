package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func getSubstring(s string, indices []int) string {
	return string(s[indices[0]:indices[1]])
}

func main() {

	arg := "strings"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	switch arg {
	case "strings":
		{
			product := "Kayak"
			fmt.Println("Contains:", strings.Contains(product, "yak"))
			fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
			fmt.Println("ContainsRune:", strings.ContainsRune(product, 'K'))
			fmt.Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
			fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
			fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))
			fmt.Println("Count:", strings.Count(product, "a"))
			fmt.Println("Index:", strings.Index(product, "a"))
			fmt.Println("LastIndex:", strings.LastIndex(product, "k"))
			fmt.Println("IndexAny:", strings.IndexAny(product, "abcdk"))
			fmt.Println("LastIndex:", strings.LastIndex(product, "k"))
			fmt.Println("LastIndexAny:", strings.LastIndexAny(product, "abcdk"))

			description := "   A boat for sailing"
			fmt.Println("Original:", strings.TrimSpace(description))

			mapper := func(r rune) rune {
				if r == 'b' {
					return 'c'
				}
				return r
			}
			mapped := strings.Map(mapper, description)
			fmt.Println("Mapped:", mapped)

			replace := strings.Replace(description, "boat", "canoe", 1)
			fmt.Println("Replaced:", strings.TrimSpace(replace))

			elements := strings.Fields(description)
			joined := strings.Join(elements, "--")
			fmt.Println("Joined:", joined)

			caser := cases.Title(language.Italian)
			fmt.Println("Title:", caser.String(description))

			caser = cases.Lower(language.Italian)
			fmt.Println("Lower:", caser.String(description))

			for _, char := range product {
				fmt.Println(string(char), "Upper case:", unicode.IsUpper(char))
			}

			isLetterB := func(r rune) bool {
				return r == 'B' || r == 'b'
			}
			fmt.Println("IndexFunc:", strings.IndexFunc(description, isLetterB))

			splits := strings.Split(description, " ")
			for _, x := range splits {
				fmt.Println("Split >>" + x + "<<")
			}
			splitsAfter := strings.SplitAfter(description, " ")
			for _, x := range splitsAfter {
				fmt.Println("SplitAfter >>" + x + "<<")
			}

			text := "It was a boat. A small boat."
			var builder strings.Builder
			for _, sub := range strings.Fields(text) {
				if sub == "small" {
					builder.WriteString("very ")
				}
				builder.WriteString(sub)
				builder.WriteRune(' ')
			}
			fmt.Println("String:", builder.String())
		}
	case "regexp":
		{
			description := "A boat for one person"
			match, err := regexp.MatchString("[A-z]oat", description)
			if err == nil {
				fmt.Println("Match:", match)
			} else {
				fmt.Println("Error:", err)
			}

			pattern, compileErr := regexp.Compile("[A-z]oat")
			question := "Is that a goat?"
			preference := "I like oats"
			if compileErr == nil {
				fmt.Println("Description:", pattern.MatchString(description))
				fmt.Println("Question:", pattern.MatchString(question))
				fmt.Println("Preference:", pattern.MatchString(preference))
			} else {
				fmt.Println("Error:", compileErr)
			}

			pattern = regexp.MustCompile("K[a-z]{4}|[A-z]oat")
			description = "Kayak. A boat for one person."
			firstIndex := pattern.FindStringIndex(description)
			allIndices := pattern.FindAllStringIndex(description, -1)
			fmt.Println("First index", firstIndex[0], "-", firstIndex[1],
				"=", getSubstring(description, firstIndex))
			for i, idx := range allIndices {
				fmt.Println("Index", i, "=", idx[0], "-",
					idx[1], "=", getSubstring(description, idx))
			}

			firstMatch := pattern.FindString(description)
			allMatches := pattern.FindAllString(description, -1)
			fmt.Println("First match:", firstMatch)
			for i, m := range allMatches {
				fmt.Println("Match", i, "=", m)
			}

			pattern = regexp.MustCompile("A ([A-z]*) for ([A-z]*)")
			subs := pattern.FindStringSubmatch(description)
			for _, s := range subs {
				fmt.Println("Match:", s)
			}

			pattern = regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
			subs = pattern.FindStringSubmatch(description)
			for _, name := range []string{"type", "capacity"} {
				fmt.Println(name, "=", subs[pattern.SubexpIndex(name)])
			}
			template := "(type: ${type}, capacity: ${capacity})"
			replaced := pattern.ReplaceAllString(description, template)
			fmt.Println(replaced)
		}
	default:
		fmt.Println("Wrong arg", arg)
	}

}
