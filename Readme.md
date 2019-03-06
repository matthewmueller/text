# text

Go Text utilities. Particularly useful for writing code generators.

## Install

```sh
go get -u github.com/matthewmueller/text
```

## API

| Method                 | From            | To               | Description |
| ---------------------- | --------------- | ---------------- | ----------- |
| `text.Base(str)`       | hi - world      | hi world         | base case   |
| `text.Lower(str)`      | TEST            | test             | lowercase   |
| `text.Camel(str)`      | Test String     | testString       | camelCase   |
| `text.Pascal(str)`     | test string     | TestString       | PascalCase  |
| `text.Snake(str)`      | Test String     | Test_String      | snake_case  |
| `text.Slug(str)`       | TestString1_2_3 | Test-String1-2-3 | slug-case   |
| `text.Dot(str)`        | Test String     | Test.String      | dot.case    |
| `text.Abbreviate(str)` | path/case       | pc               | a           |
| `text.Singular(str)`   | My-Entrées      | My-Entrée        | singular    |
| `text.Plural(str)`     | My-Entrées      | My-Entrée        | plurals     |

## Thanks

- https://github.com/blakeembrey/change-case
- https://github.com/ianstormtaylor/to-case

## License

MIT
