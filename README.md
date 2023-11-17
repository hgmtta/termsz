# Termsz
**Termsz** is a cross-platform library for getting the size of the current terminal window.

## Installation
To install the latest version of the library, run the following command:
```
go get -u github.com/hgmtta/termsz@latest
```

## Usage
First, include the library in your project:
```go
import "github.com/hgmtta/termsz"
```

Then, call the `termsz.GetSize()` function to get the size of the current terminal window:
```go
cols, rows, err := termsz.GetSize()
if err != nil {
   panic(err)
}
fmt.Printf("Columns: %d, Rows: %d\n", cols, rows)
```
For convenience, you can also use the `termsz.GetColumns()` and `termsz.GetRows()` functions to get the number of columns and rows, respectively.

For more details, see the [documentation](https://pkg.go.dev/github.com/hgmtta/termsz).

## License
This library is licensed under the MIT License. See [LICENSE.md](LICENSE.md) for more details.