package termsz

// GetSize returns the size of the current terminal window.
// It returns the number of columns and rows as integers and an error if any occurs.
func GetSize() (cols, rows int, err error) {
	return getSize()
}

// GetColumns returns the number of columns in the current terminal window.
// It returns the number of columns as an integer and an error if any occurs.
func GetColumns() (int, error) {
	cols, _, err := GetSize()
	return cols, err
}
