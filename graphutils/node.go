package graphutils

import "github.com/goccy/go-graphviz/cgraph"

// NewNode creates a named cgraph.Node. Panics on error.
func NewNode(graph *cgraph.Graph, name string) *cgraph.Node {
  result, err := graph.CreateNode(name)
  if err != nil {
    panic(err)
  }

  return result
}
