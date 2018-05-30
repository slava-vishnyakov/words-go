package words

import (
	"strings"
	"unicode"
	"regexp"
)

var chineseRegexp = regexp.MustCompile("[\u4e00-\u9fa5]")

var italianPrefix = map[string]bool{
	"d":     true,
	"c":     true,
	"l":     true,
	"cos":   true,
	"dov":   true,
	"anch":  true,
	"un":    true,
	"com":   true,
	"quant": true,
	"dev":   true,
}

var englishEndings = map[string]bool{
	"t":  true,
	"s":  true,
	"m":  true,
	"ll": true,
	"d":  true,
	"ve": true,
	"re": true,
}

func nonWord(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsDigit(r)
}

func Words(text string) []string {
	var words []string

	// return runes for chinese
	if chineseRegexp.MatchString(text) {
		runes := []rune(text)
		for i := range runes {
			if !unicode.IsSpace(runes[i]) && !unicode.IsPunct(runes[i]) {
				words = append(words, string(runes[i]))
			}
		}

		return words
	}

	words = strings.FieldsFunc(strings.ToLower(text), nonWord)

	res := []string{}
	for i := 0; i < len(words); i++ {
		if len(res) > 0 {
			if englishEndings[words[i]] {
				res[len(res)-1] = res[len(res)-1] + "'" + words[i]
				continue
			}
			if italianPrefix[res[len(res)-1]] {
				res[len(res)-1] = res[len(res)-1] + "'" + words[i]
				continue
			}
		}
		res = append(res, words[i])
	}

	return res
}

func RemoveWords(s string, words []string, repl string) string {
	for _, w := range words {
		r := regexp.MustCompile(`(^|\P{L})\Q` + w + `\E($|\P{L})`)
		s = r.ReplaceAllString(s, `$1`+repl+`$2`)
	}
	r2 := regexp.MustCompile(`(\Q` + repl + `\E)(\P{L}+\Q` + repl + `\E)+`)
	s = r2.ReplaceAllString(s, `$1`)
	return strings.TrimFunc(s, nonWord)
}
