package text

import (
	"strings"
	"unicode"

	"github.com/gedex/inflector"
)

// Base case
func Base(s string) string {
	in := []rune(s)
	out := []rune{}

	for i := 0; i < len(in); i++ {
		r := in[i]

		// ignore certain punctuation
		if r == '\'' {
			continue
		}

		// camelCase -> camel case
		if i >= 1 && unicode.IsUpper(r) && unicode.IsLower(in[i-1]) {
			out = append(out, ' ')
			out = append(out, r)
			continue
		}

		// 123Case -> 123 case
		if i >= 1 && unicode.IsUpper(r) && unicode.IsDigit(in[i-1]) {
			out = append(out, ' ')
			out = append(out, r)
			continue
		}

		// CAMELCase -> camel case
		if i >= 1 && i < len(in)-1 && unicode.IsUpper(in[i-1]) && unicode.IsUpper(r) && unicode.IsLower(in[i+1]) {
			out = append(out, ' ')
			out = append(out, r)
			continue
		}

		// letters to lowercase
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			out = append(out, r)
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
	a := strings.Split(Lower(Base(s)), " ")
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
	a := strings.Split(Lower(Base(s)), " ")
	for i := range a {
		a[i] = strings.Title(a[i])
	}
	return strings.Join(a, "")
}

// Snake fn
func Snake(s string) string {
	a := strings.Split(Base(s), " ")
	return strings.Join(a, "_")
}

// Slug case
func Slug(s string) string {
	a := strings.Split(Base(s), " ")
	return strings.Join(a, "-")
}

// Title case a string
func Title(s string) string {
	return strings.Title(s)
}

// Dot case
// e.g. "Test String" => "test.string"
func Dot(s string) string {
	a := strings.Split(Base(s), " ")
	return strings.Join(a, ".")
}

// Abbreviate case
func Abbreviate(s string) string {
	a := strings.Split(Base(s), " ")
	o := ""
	for _, w := range a {
		if len(w) == 0 {
			continue
		}
		o += string(w[0])
	}
	return o
}

// Singular string
func Singular(s string) string {
	a := strings.Split(Base(s), " ")
	last := a[len(a)-1]
	i := strings.LastIndex(s, last)

	if i < -1 {
		return inflector.Singularize(s)
	}

	o := s[:i]
	o += inflector.Singularize(last)
	o += s[i+len(last):]

	return o
}

// Plural string
func Plural(s string) string {
	a := strings.Split(Base(s), " ")
	last := a[len(a)-1]
	i := strings.LastIndex(s, last)

	if i < -1 {
		return inflector.Pluralize(s)
	}

	o := s[:i]
	o += inflector.Pluralize(last)
	o += s[i+len(last):]

	return o
}
