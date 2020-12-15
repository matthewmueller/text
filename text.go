package text

import (
	"strings"
	"unicode"

	"github.com/gedex/inflector"
)

// Space case (e.g. Space case)
func Space(s string) string {
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

// Lower case is just an alias to strings.ToLower
var Lower = strings.ToLower

// Upper case is just an alias to strings.ToUpper
var Upper = strings.ToUpper

// Title case (e.g. Title Case)
func Title(s string) string {
	a := strings.Split(Lower(Space(s)), " ")
	for i := range a {
		a[i] = strings.Title(a[i])
	}
	return strings.Join(a, " ")
}

// Camel case (e.g. camelCase)
func Camel(s string) string {
	a := strings.Split(Lower(Space(s)), " ")
	for i := range a {
		if i == 0 {
			continue
		}
		a[i] = strings.Title(a[i])
	}
	return strings.Join(a, "")
}

// Pascal case (e.g. PascalCase)
func Pascal(s string) string {
	a := strings.Split(Lower(Space(s)), " ")
	for i := range a {
		a[i] = strings.Title(a[i])
	}
	return strings.Join(a, "")
}

// Snake case (e.g. snake_case)
func Snake(s string) string {
	a := strings.Split(Space(s), " ")
	return strings.Join(a, "_")
}

// Slug case (e.g. snake-case)
func Slug(s string) string {
	a := strings.Split(Space(s), " ")
	return strings.Join(a, "-")
}

// Dot case (e.g. dot.case)
func Dot(s string) string {
	a := strings.Split(Space(s), " ")
	return strings.Join(a, ".")
}

// Path case (e.g. path/case)
func Path(s string) string {
	a := strings.Split(Space(s), " ")
	return strings.Join(a, "/")
}

// Short case (e.g. sc)
func Short(s string) string {
	a := strings.Split(Space(s), " ")
	o := ""
	for _, w := range a {
		if len(w) == 0 {
			continue
		}
		o += string(w[0])
	}
	return o
}

// Slim case (e.g. Slimcase)
func Slim(s string) string {
	a := strings.Split(Space(s), " ")
	return strings.Join(a, "")
}

// Singular string (e.g. apple)
func Singular(s string) string {
	a := strings.Split(Space(s), " ")
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

// Plural string (e.g. apples)
func Plural(s string) string {
	a := strings.Split(Space(s), " ")
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
