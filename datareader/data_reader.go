package datareader

import (
	"encoding/csv"
	"os"
)

// ReadData reads the contents from a CSV filepath. A CSV filepath is assumed to be supplied. Panics on error.
func ReadData(filepath string) [][]string {
	file, openFileError := os.Open(filepath)
	if openFileError != nil {
		panic(openFileError)
	}
	defer func(file *os.File) {
		closeFileError := file.Close()
		if closeFileError != nil {
			panic(closeFileError)
		}
	}(file)

	reader := csv.NewReader(file)
	data, readAllError := reader.ReadAll()
	if readAllError != nil {
		panic(readAllError)
	}

	return data
}
