package text_test

import (
	"testing"

	text "github.com/matthewmueller/go-text"
	"github.com/stretchr/testify/assert"
)

type test struct {
	fn  func(string) string
	in  string
	out string
}

func TestNoCase(t *testing.T) {
	tests := []test{
		{text.No, "test", "test"},
		{text.No, "TEST", "test"},
		{text.No, "testString", "test string"},
		{text.No, "testString123", "test string123"},
		{text.No, "hi-world", "hi world"},
		{text.No, "hi - world", "hi world"},
		{text.No, "hi__world", "hi world"},
		{text.No, "hi_-_world", "hi world"},
		{text.No, "testString_1_2_3", "test string 1 2 3"},
		{text.No, "x_256", "x 256"},
		{text.No, "anHTMLTag", "an html tag"},
		{text.No, "ID123String", "id123 string"},
		{text.No, "Id123String", "id123 string"},
		{text.No, "foo bar123", "foo bar123"},
		{text.No, "a1bStar", "a1b star"},
		{text.No, "CONSTANT_CASE ", "constant case"},
		{text.No, "CONST123_FOO", "const123 foo"},
		{text.No, "FOO_bar", "foo bar"},
		{text.No, "dot.case", "dot case"},
		{text.No, "path/case", "path case"},
		{text.No, "snake_case", "snake case"},
		{text.No, "snake_case123", "snake case123"},
		{text.No, "snake_case_123", "snake case 123"},
		{text.No, "\"quotes\"", "quotes"},
		{text.No, "version 0.45.0", "version 0 45 0"},
		{text.No, "version 0..78..9", "version 0 78 9"},
		{text.No, "version 4_99/4", "version 4 99 4"},
		{text.No, "  test  ", "test"},
		{text.No, "  te+st  ", "te st"},
		{text.No, "hi-world", "hi world"},
		{text.No, "hi+world", "hi world"},
		{text.No, "hi\\world", "hi world"},
		{text.No, "hi/world", "hi world"},
		{text.No, "hi*world", "hi world"},
		{text.No, "hi#world", "hi world"},
		{text.No, "hi&world", "hi world"},
		{text.No, "español", "español"},
		{text.No, "Beyoncé Knowles", "beyoncé knowles"},
		{text.No, "Iñtërnâtiônàlizætiøn", "iñtërnâtiônàlizætiøn"},
		{text.No, "something_2014_other", "something 2014 other"},
		{text.No, "HELLO WORLD!", "hello world"},
		{text.No, "foo bar!", "foo bar"},
		{text.No, "A STRING", "a string"},
		{text.No, "amazon s3 data", "amazon s3 data"},
		{text.No, "foo_13_bar", "foo 13 bar"},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestLowerCase(t *testing.T) {
	tests := []test{
		{text.Lower, "TEST", "test"},
		{text.Lower, "test", "test"},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in))
	}
}

func TestCamelCase(t *testing.T) {
	tests := []test{
		{text.Camel, "test", "test"},
		{text.Camel, "TEST", "test"},
		{text.Camel, "test string", "testString"},
		{text.Camel, "Test String", "testString"},
		{text.Camel, "dot.case", "dotCase"},
		{text.Camel, "path/case", "pathCase"},
		{text.Camel, "version 1.2.10", "version1210"},
		{text.Camel, "version 1.21.0", "version1210"},
		{text.Camel, "TestString", "testString"},
		{text.Camel, "simple éxample", "simpleÉxample"},
		{text.Camel, "test 1 2 3", "test123"},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in))
	}
}

func TestPascalCase(t *testing.T) {
	tests := []test{
		{text.Pascal, "test", "Test"},
		{text.Pascal, "TEST", "Test"},
		{text.Pascal, "test string", "TestString"},
		{text.Pascal, "Test String", "TestString"},
		{text.Pascal, "dot.case", "DotCase"},
		{text.Pascal, "path/case", "PathCase"},
		{text.Pascal, "TestString", "TestString"},
		{text.Pascal, "test 1 2 3", "Test123"},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in))
	}
}

func TestSnakeCase(t *testing.T) {
	tests := []test{
		{text.Snake, "test", "test"},
		{text.Snake, "TEST", "test"},
		{text.Snake, "test string", "test_string"},
		{text.Snake, "Test String", "test_string"},
		{text.Snake, "dot.case", "dot_case"},
		{text.Snake, "path/case", "path_case"},
		{text.Snake, "TestString", "test_string"},
		{text.Snake, "TestString1_2_3", "test_string1_2_3"},
		{text.Snake, "My Entrée", "my_entrée"},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in))
	}
}

func TestAbbrCase(t *testing.T) {
	tests := []test{
		{text.Abbreviation, "test", "t"},
		{text.Abbreviation, "TEST", "t"},
		{text.Abbreviation, "test string", "ts"},
		{text.Abbreviation, "Test String", "ts"},
		{text.Abbreviation, "dot.case", "dc"},
		{text.Abbreviation, "path/case", "pc"},
		{text.Abbreviation, "TestString", "ts"},
		{text.Abbreviation, "TestString1_2_3", "ts23"},
		{text.Abbreviation, "My Entrée", "me"},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in))
	}
}
