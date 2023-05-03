package datareader

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

// Associate represents an associate model. The RootID is the identifier of the parent node.
// ID is the identifier of this node.
type Associate struct {
	Name   string
	RootID string
	ID     string
}

const (
	_header = 0
)

// AssociatesFromCSVData returns a list of Associate from the data matrix. Assumes the Associate fields are in the
// following positions (zero-indexed):
// Associate.Name = 4, Associate.RootID = 1, and Associate.ID = 2.
// Panics if assumed indexes are violated.
func AssociatesFromCSVData(data [][]string) (results []Associate) {
	for i, row := range data {
		if i == _header {
			continue
		}

		associate := Associate{
			Name:   cases.Title(language.Spanish).String(strings.ToLower(row[4])),
			RootID: row[1],
			ID:     row[2],
		}

		results = append(results, associate)
	}

	return results
}
