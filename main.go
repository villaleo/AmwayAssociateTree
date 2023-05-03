package main

import (
  "amway-associate-tree/datareader"
  "amway-associate-tree/graphutils"
  "github.com/goccy/go-graphviz"
  "github.com/goccy/go-graphviz/cgraph"
  "log"
  "os"
  "path"
  "regexp"
  "strings"
)

func main() {
  filePath, outputFileName := readCommandLineArgs()

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

  log.Println("Generating graph...")
  createTreeFromAssociatesFile(graph, filePath)

  graph.Set("fontname", "Helvetica,Arial,sans-serif")
  graph.SetBackgroundColor("lightblue1")
  graph.SetLayout("dot")

  err = visualGraph.RenderFilename(graph, graphviz.SVG, outputFileName+".svg")
  if err != nil {
    panic(err)
    return
  }

  log.Println("Done!")
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

func readCommandLineArgs() (filePath, outputFileName string) {
  args := os.Args[1:]
  if argsLength := len(args); argsLength == 0 {
    showUsage()

  } else if argsLength == 1 {
    if isCSVFile, _ := path.Match("**/*.csv", prefixIfInCWD(args[0])); !isCSVFile {
      log.Fatalln("Expected a CSV file for input data.")
    }

    filePath = args[0]
    outputFileName = "results"
    return

  } else if argsLength == 2 {
    if isCSVFile, _ := path.Match("**/*.csv", prefixIfInCWD(args[0])); !isCSVFile {
      log.Fatalln("Expected a CSV file for input data.")
    }
    if args[1] != "-o" {
      log.Fatalf("Expected '-o', got '%s'.\n", args[1])
    }
    log.Fatalln("Missing required positional argument: <output-filename>")

  } else if argsLength == 3 {
    if isCSVFile, _ := path.Match("**/*.csv", prefixIfInCWD(args[0])); !isCSVFile {
      log.Fatalln("Expected a CSV file for input data.")
    }
    if args[1] != "-o" {
      log.Fatalf("Expected '-o', got '%s'.\n", args[1])
    }
    if isValidFileName, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, args[2]); !isValidFileName {
      log.Fatal(
        "Invalid output-filename. Filename must be alphanumeric, " +
          "with '-' and '_' being the only valid symbols.",
      )
    }

    filePath = args[0]
    outputFileName = args[2]
    return

  }

  log.Fatalln("Excess arguments supplied.")
  return
}

func showUsage() {
  log.Fatalln("Usage: go run . <filepath> [-o <output-filename = results.svg]")
}

func prefixIfInCWD(filePath string) string {
  if filePath[0] != '/' && strings.Count(filePath, "/") == 0 {
    return "/" + filePath
  }
  return filePath
}
