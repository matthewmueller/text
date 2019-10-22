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
		{text.Base, "EThreader", "E Threader"},
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

func TestDotCase(t *testing.T) {
	tests := []test{
		{text.Dot, "test", "test"},
		{text.Dot, "TEST", "TEST"},
		{text.Dot, "test string", "test.string"},
		{text.Dot, "Test String", "Test.String"},
		{text.Dot, "dot.case", "dot.case"},
		{text.Dot, "path/case", "path.case"},
		{text.Dot, "TestString", "Test.String"},
		{text.Dot, "TestString1_2_3", "Test.String1.2.3"},
		{text.Dot, "TestString_1_2_3", "Test.String.1.2.3"},
		{text.Dot, "My Entrée", "My.Entrée"},
		{text.Dot, "MY STRING", "MY.STRING"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestAbbreviateCase(t *testing.T) {
	tests := []test{
		{text.Abbreviate, "", ""},
		{text.Abbreviate, "test", "t"},
		{text.Abbreviate, "TEST", "T"},
		{text.Abbreviate, "test string", "ts"},
		{text.Abbreviate, "test String", "tS"},
		{text.Abbreviate, "Test String", "TS"},
		{text.Abbreviate, "dot.case", "dc"},
		{text.Abbreviate, "path/case", "pc"},
		{text.Abbreviate, "TestString", "TS"},
		{text.Abbreviate, "TestString1_2_3", "TS23"},
		{text.Abbreviate, "My Entrée", "ME"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestSingular(t *testing.T) {
	tests := []test{
		{text.Singular, "", ""},
		{text.Singular, "tests", "test"},
		{text.Singular, "TESTS", "TEST"},
		{text.Singular, "test strings", "test string"},
		{text.Singular, "Test Strings", "Test String"},
		{text.Singular, "dot.cases", "dot.case"},
		{text.Singular, "paths/cases", "paths/case"},
		{text.Singular, "TestsStrings", "TestsString"},
		{text.Singular, "TestString1_2_3s", "TestString1_2_3"},
		{text.Singular, "My-Entrées", "My-Entrée"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestPlural(t *testing.T) {
	tests := []test{
		{text.Plural, "", ""},
		{text.Plural, "test", "tests"},
		{text.Plural, "TEST", "TESTs"},
		{text.Plural, "test string", "test strings"},
		{text.Plural, "Test String", "Test Strings"},
		{text.Plural, "dot.case", "dot.cases"},
		{text.Plural, "path/case", "path/cases"},
		{text.Plural, "TestString", "TestStrings"},
		{text.Plural, "TestString1_2_3", "TestString1_2_3s"},
		{text.Plural, "My Entrée", "My Entrées"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}
