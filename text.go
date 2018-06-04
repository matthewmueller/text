package text

import (
	"strings"
	"unicode"
)

// No clears the case
func No(s string) string {
	in := []rune(s)
	out := []rune{}

	for i := 0; i < len(in); i++ {
		r := in[i]

		// camelCase -> camel case
		if i >= 1 && unicode.IsUpper(r) && unicode.IsLower(in[i-1]) {
			out = append(out, ' ')
			out = append(out, unicode.ToLower(r))
			continue
		}

		// 123Case -> 123 case
		if i >= 1 && unicode.IsUpper(r) && unicode.IsDigit(in[i-1]) {
			out = append(out, ' ')
			out = append(out, unicode.ToLower(r))
			continue
		}

		// CAMELCase -> camel case
		if i >= 1 && i < len(in)-1 && unicode.IsUpper(in[i-1]) && unicode.IsUpper(r) && unicode.IsLower(in[i+1]) && unicode.IsUpper(in[i-2]) {
			out = append(out, ' ')
			out = append(out, unicode.ToLower(r))
			continue
		}

		// letters to lowercase
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			out = append(out, unicode.ToLower(r))
			continue
		}

		// punctuation and spaces get trimmed into a single space
		if unicode.IsPunct(r) || unicode.IsSpace(r) || unicode.IsSymbol(r) {
			if i >= 1 && !unicode.IsPunct(in[i-1]) && !unicode.IsSpace(in[i-1]) && !unicode.IsSymbol(in[i-1]) {
				out = append(out, ' ')
				continue
			}
		}
	}

	// TODO: can we trim without going back through the string again?
	return strings.TrimSpace(string(out))
}

// Lower fn
func Lower(s string) string {
	return strings.ToLower(s)
}

// Upper fn
func Upper(s string) string {
	return strings.ToUpper(s)
}

// Camel fn
func Camel(s string) string {
	a := strings.Split(No(s), " ")
	for i := range a {
		if i == 0 {
			continue
		}
		a[i] = strings.Title(a[i])
	}
	return strings.Join(a, "")
}

// Pascal fn
func Pascal(s string) string {
	a := strings.Split(No(s), " ")
	for i := range a {
		a[i] = strings.Title(a[i])
	}
	return strings.Join(a, "")
}

// Snake fn
func Snake(s string) string {
	a := strings.Split(No(s), " ")
	return strings.Join(a, "_")
}

// Title case a string
func Title(s string) string {
	return strings.Title(s)
}

// Abbreviation case
func Abbreviation(s string) string {
	a := strings.Split(No(s), " ")
	o := make([]byte, len(a))
	for i := range a {
		o[i] = a[i][0]
	}
	return string(o)
}

// func Dot(s string) string {
// return s
// }

// func Swap(s string) string {

// }

// func Path(s string) string {

// }

// Param fn
// func Param(s string) string {

// }

// Header fn
// func Header(s string) string {

// }

// Constant fn
// func Constant(s string) string {

// }

// Sentence fn
// func Sentence(s string) string {

// }
