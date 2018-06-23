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

func TestBaseCase(t *testing.T) {
	tests := []test{
		{text.Base, "", ""},
		{text.Base, "test", "test"},
		{text.Base, "TEST", "TEST"},
		{text.Base, "Anki's Trip", "Ankis Trip"},
		{text.Base, "testString", "test String"},
		{text.Base, "testString123", "test String123"},
		{text.Base, "hi-world", "hi world"},
		{text.Base, "hi - world", "hi world"},
		{text.Base, "hi__world", "hi world"},
		{text.Base, "hi_-_world", "hi world"},
		{text.Base, "testString_1_2_3", "test String 1 2 3"},
		{text.Base, "x_256", "x 256"},
		{text.Base, "anHTMLTag", "an HTML Tag"},
		{text.Base, "ID123String", "ID123 String"},
		{text.Base, "Id123String", "Id123 String"},
		{text.Base, "foo bar123", "foo bar123"},
		{text.Base, "a1bStar", "a1b Star"},
		{text.Base, "CONSTANT_CASE ", "CONSTANT CASE"},
		{text.Base, "CONST123_FOO", "CONST123 FOO"},
		{text.Base, "FOO_bar", "FOO bar"},
		{text.Base, "dot.case", "dot case"},
		{text.Base, "path/case", "path case"},
		{text.Base, "snake_case", "snake case"},
		{text.Base, "snake_case123", "snake case123"},
		{text.Base, "snake_case_123", "snake case 123"},
		{text.Base, "\"quotes\"", "quotes"},
		{text.Base, "version 0.45.0", "version 0 45 0"},
		{text.Base, "version 0..78..9", "version 0 78 9"},
		{text.Base, "version 4_99/4", "version 4 99 4"},
		{text.Base, "  test  ", "test"},
		{text.Base, "  te+st  ", "te st"},
		{text.Base, "hi-world", "hi world"},
		{text.Base, "hi+world", "hi world"},
		{text.Base, "hi\\world", "hi world"},
		{text.Base, "hi/world", "hi world"},
		{text.Base, "hi*world", "hi world"},
		{text.Base, "hi#world", "hi world"},
		{text.Base, "hi&world", "hi world"},
		{text.Base, "español", "español"},
		{text.Base, "Beyoncé Knowles", "Beyoncé Knowles"},
		{text.Base, "Iñtërnâtiônàlizætiøn", "Iñtërnâtiônàlizætiøn"},
		{text.Base, "something_2014_other", "something 2014 other"},
		{text.Base, "HELLO WORLD!", "HELLO WORLD"},
		{text.Base, "foo bar!", "foo bar"},
		{text.Base, "A STRING", "A STRING"},
		{text.Base, "amazon s3 data", "amazon s3 data"},
		{text.Base, "foo_13_bar", "foo 13 bar"},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestLowerCase(t *testing.T) {
	tests := []test{
		{text.Lower, "", ""},
		{text.Lower, "TEST", "test"},
		{text.Lower, "test", "test"},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestCamelCase(t *testing.T) {
	tests := []test{
		{text.Camel, "", ""},
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
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestPascalCase(t *testing.T) {
	tests := []test{
		{text.Pascal, "", ""},
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
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestSnakeCase(t *testing.T) {
	tests := []test{
		{text.Snake, "", ""},
		{text.Snake, "test", "test"},
		{text.Snake, "TEST", "TEST"},
		{text.Snake, "test string", "test_string"},
		{text.Snake, "Test String", "Test_String"},
		{text.Snake, "dot.case", "dot_case"},
		{text.Snake, "path/case", "path_case"},
		{text.Snake, "TestString", "Test_String"},
		{text.Snake, "TestString1_2_3", "Test_String1_2_3"},
		{text.Snake, "My Entrée", "My_Entrée"},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestSlugCase(t *testing.T) {
	tests := []test{
		{text.Slug, "", ""},
		{text.Slug, "test", "test"},
		{text.Slug, "TEST", "TEST"},
		{text.Slug, "test string", "test-string"},
		{text.Slug, "Test String", "Test-String"},
		{text.Slug, "dot.case", "dot-case"},
		{text.Slug, "path/case", "path-case"},
		{text.Slug, "TestString", "Test-String"},
		{text.Slug, "TestString1_2_3", "Test-String1-2-3"},
		{text.Slug, "My Entrée", "My-Entrée"},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestAbbrCase(t *testing.T) {
	tests := []test{
		{text.Abbreviation, "", ""},
		{text.Abbreviation, "test", "t"},
		{text.Abbreviation, "TEST", "T"},
		{text.Abbreviation, "test string", "ts"},
		{text.Abbreviation, "test String", "tS"},
		{text.Abbreviation, "Test String", "TS"},
		{text.Abbreviation, "dot.case", "dc"},
		{text.Abbreviation, "path/case", "pc"},
		{text.Abbreviation, "TestString", "TS"},
		{text.Abbreviation, "TestString1_2_3", "TS23"},
		{text.Abbreviation, "My Entrée", "ME"},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}
