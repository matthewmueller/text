package text_test

import (
	"testing"

	"github.com/matthewmueller/text"
	"github.com/stretchr/testify/assert"
)

type test struct {
	fn  func(string) string
	in  string
	out string
}

func TestSpaceCase(t *testing.T) {
	tests := []test{
		{text.Space, "", ""},
		{text.Space, "test", "test"},
		{text.Space, "TEST", "TEST"},
		{text.Space, "Anki's Trip", "Ankis Trip"},
		{text.Space, "testString", "test String"},
		{text.Space, "testString123", "test String123"},
		{text.Space, "hi-world", "hi world"},
		{text.Space, "hi - world", "hi world"},
		{text.Space, "hi__world", "hi world"},
		{text.Space, "hi_-_world", "hi world"},
		{text.Space, "testString_1_2_3", "test String 1 2 3"},
		{text.Space, "x_256", "x 256"},
		{text.Space, "anHTMLTag", "an HTML Tag"},
		{text.Space, "ID123String", "ID123 String"},
		{text.Space, "Id123String", "Id123 String"},
		{text.Space, "foo bar123", "foo bar123"},
		{text.Space, "a1bStar", "a1b Star"},
		{text.Space, "CONSTANT_CASE ", "CONSTANT CASE"},
		{text.Space, "CONST123_FOO", "CONST123 FOO"},
		{text.Space, "FOO_bar", "FOO bar"},
		{text.Space, "dot.case", "dot case"},
		{text.Space, "path/case", "path case"},
		{text.Space, "snake_case", "snake case"},
		{text.Space, "snake_case123", "snake case123"},
		{text.Space, "snake_case_123", "snake case 123"},
		{text.Space, "\"quotes\"", "quotes"},
		{text.Space, "version 0.45.0", "version 0 45 0"},
		{text.Space, "version 0..78..9", "version 0 78 9"},
		{text.Space, "version 4_99/4", "version 4 99 4"},
		{text.Space, "  test  ", "test"},
		{text.Space, "  te+st  ", "te st"},
		{text.Space, "hi-world", "hi world"},
		{text.Space, "hi+world", "hi world"},
		{text.Space, "hi\\world", "hi world"},
		{text.Space, "hi/world", "hi world"},
		{text.Space, "hi*world", "hi world"},
		{text.Space, "hi#world", "hi world"},
		{text.Space, "hi&world", "hi world"},
		{text.Space, "español", "español"},
		{text.Space, "Beyoncé Knowles", "Beyoncé Knowles"},
		{text.Space, "Iñtërnâtiônàlizætiøn", "Iñtërnâtiônàlizætiøn"},
		{text.Space, "something_2014_other", "something 2014 other"},
		{text.Space, "HELLO WORLD!", "HELLO WORLD"},
		{text.Space, "foo bar!", "foo bar"},
		{text.Space, "A STRING", "A STRING"},
		{text.Space, "amazon s3 data", "amazon s3 data"},
		{text.Space, "foo_13_bar", "foo 13 bar"},
		{text.Space, "EThreader", "E Threader"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestLowerCase(t *testing.T) {
	tests := []test{
		{text.Lower, "", ""},
		{text.Lower, "test", "test"},
		{text.Lower, "TEST", "test"},
		{text.Lower, "test string", "test string"},
		{text.Lower, "Test String", "test string"},
		{text.Lower, "dot.case", "dot.case"},
		{text.Lower, "path/case", "path/case"},
		{text.Lower, "version 1.2.10", "version 1.2.10"},
		{text.Lower, "version 1.21.0", "version 1.21.0"},
		{text.Lower, "TestString", "teststring"},
		{text.Lower, "simple éxample", "simple éxample"},
		{text.Lower, "test 1 2 3", "test 1 2 3"},
		{text.Lower, "Out0", "out0"},
		{text.Lower, "out0", "out0"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}
func TestUpperCase(t *testing.T) {
	tests := []test{
		{text.Upper, "", ""},
		{text.Upper, "test", "TEST"},
		{text.Upper, "TEST", "TEST"},
		{text.Upper, "test string", "TEST STRING"},
		{text.Upper, "Test String", "TEST STRING"},
		{text.Upper, "dot.case", "DOT.CASE"},
		{text.Upper, "path/case", "PATH/CASE"},
		{text.Upper, "version 1.2.10", "VERSION 1.2.10"},
		{text.Upper, "version 1.21.0", "VERSION 1.21.0"},
		{text.Upper, "TestString", "TESTSTRING"},
		{text.Upper, "simple éxample", "SIMPLE ÉXAMPLE"},
		{text.Upper, "test 1 2 3", "TEST 1 2 3"},
		{text.Upper, "Out0", "OUT0"},
		{text.Upper, "out0", "OUT0"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}
func TestTitleCase(t *testing.T) {
	tests := []test{
		{text.Title, "", ""},
		{text.Title, "test", "Test"},
		{text.Title, "TEST", "Test"},
		{text.Title, "test string", "Test String"},
		{text.Title, "Test String", "Test String"},
		{text.Title, "dot.case", "Dot Case"},
		{text.Title, "path/case", "Path Case"},
		{text.Title, "version 1.2.10", "Version 1 2 10"},
		{text.Title, "version 1.21.0", "Version 1 21 0"},
		{text.Title, "TestString", "Test String"},
		{text.Title, "simple éxample", "Simple Éxample"},
		{text.Title, "test 1 2 3", "Test 1 2 3"},
		{text.Title, "Out0", "Out0"},
		{text.Title, "out0", "Out0"},
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
		{text.Camel, "Out0", "out0"},
		{text.Camel, "out0", "out0"},
		{text.Camel, "Blog", "blog"},
		{text.Camel, "ASKED", "asked"},
		{text.Camel, "ASKED_RECIEVED", "askedRecieved"},
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
		{text.Pascal, "Out0", "Out0"},
		{text.Pascal, "out0", "Out0"},
		{text.Pascal, "Blog", "Blog"},
		{text.Pascal, "ASKED", "Asked"},
		{text.Pascal, "ASKED_RECIEVED", "AskedRecieved"},
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
		{text.Snake, "Out0", "Out0"},
		{text.Snake, "out0", "out0"},
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
		{text.Slug, "Out0", "Out0"},
		{text.Slug, "out0", "out0"},
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
		{text.Dot, "Out0", "Out0"},
		{text.Dot, "out0", "out0"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestPathCase(t *testing.T) {
	tests := []test{
		{text.Path, "test", "test"},
		{text.Path, "TEST", "TEST"},
		{text.Path, "test string", "test/string"},
		{text.Path, "Test String", "Test/String"},
		{text.Path, "dot.case", "dot/case"},
		{text.Path, "path/case", "path/case"},
		{text.Path, "TestString", "Test/String"},
		{text.Path, "TestString1_2_3", "Test/String1/2/3"},
		{text.Path, "TestString_1_2_3", "Test/String/1/2/3"},
		{text.Path, "My Entrée", "My/Entrée"},
		{text.Path, "MY STRING", "MY/STRING"},
		{text.Path, "Out0", "Out0"},
		{text.Path, "out0", "out0"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestShortCase(t *testing.T) {
	tests := []test{
		{text.Short, "", ""},
		{text.Short, "test", "t"},
		{text.Short, "TEST", "T"},
		{text.Short, "test string", "ts"},
		{text.Short, "test String", "tS"},
		{text.Short, "Test String", "TS"},
		{text.Short, "dot.case", "dc"},
		{text.Short, "path/case", "pc"},
		{text.Short, "TestString", "TS"},
		{text.Short, "TestString1_2_3", "TS23"},
		{text.Short, "My Entrée", "ME"},
		{text.Short, "Out0", "O"},
		{text.Short, "out0", "o"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestSlimCase(t *testing.T) {
	tests := []test{
		{text.Slim, "", ""},
		{text.Slim, "test", "test"},
		{text.Slim, "TEST", "TEST"},
		{text.Slim, "test string", "teststring"},
		{text.Slim, "test String", "testString"},
		{text.Slim, "Test String", "TestString"},
		{text.Slim, "dot.case", "dotcase"},
		{text.Slim, "path/case", "pathcase"},
		{text.Slim, "TestString", "TestString"},
		{text.Slim, "TestString1_2_3", "TestString123"},
		{text.Slim, "My Entrée", "MyEntrée"},
		{text.Slim, "Out0", "Out0"},
		{text.Slim, "out0", "out0"},
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
		{text.Singular, "Out0", "Out0"},
		{text.Singular, "out0", "out0"},
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
		{text.Plural, "Out0", "Out0s"},
		{text.Plural, "out0", "out0s"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}
