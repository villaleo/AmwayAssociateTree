package main

import (
	"amway-associate-tree/datareader"
	"amway-associate-tree/graphutils"
	"fmt"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

func main() {
	visualGraph := graphviz.New()
	defer func(g *graphviz.Graphviz) {
		err := g.Close()
		if err != nil {
			panic(err)
		}
	}(visualGraph)

	graph, err := visualGraph.Graph()
	if err != nil {
		panic(err)
		return
	}

	fmt.Println("Generating graph...")
	createTreeFromAssociatesFile(graph, "associates.csv")

	graph.Set("fontname", "Helvetica,Arial,sans-serif")
	graph.SetBackgroundColor("lightblue1")
	graph.SetLayout("dot")

	err = visualGraph.RenderFilename(graph, graphviz.SVG, "amway_associate_tree.svg")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Done!")
}

func createTreeFromAssociatesFile(graph *cgraph.Graph, filepath string) {
	dataMatrix := datareader.ReadData(filepath)
	associates := datareader.AssociatesFromCSVData(dataMatrix)

	nodes := make(map[string]*cgraph.Node)
	for _, associate := range associates {
		node := graphutils.NewNode(graph, associate.Name)
		nodes[associate.ID] = node
		styleNode(node)
	}

	for _, associate := range associates {
		parent := nodes[associate.RootID]
		if parent == nil {
			nodes[associate.ID].SetFillColor("lightblue4").SetFontColor("white")
			continue
		}
		graphutils.NewEdge(graph, parent, nodes[associate.ID])
	}
}

func styleNode(node *cgraph.Node) {
	node.SetShape(cgraph.RectShape)
	node.SetStyle(cgraph.FilledNodeStyle)
	node.SetFillColor("lightblue3")
	node.SetFontSize(10)
}
