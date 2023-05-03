package graphutils

import "github.com/goccy/go-graphviz/cgraph"

// NewEdge creates an edge from one cgraph.Node to another. Panics on error.
func NewEdge(graph *cgraph.Graph, from, to *cgraph.Node) *cgraph.Edge {
  result, err := graph.CreateEdge("", from, to)
  if err != nil {
    panic(err)
  }

  return result
}
