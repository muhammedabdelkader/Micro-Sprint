package main

import (
	"fmt"

	"github.com/awalterschulze/gographviz"
)

/*
This module creates a new graph object and sets the name and direction of the graph. Then it creates subgraphs for each threat type in the STRIDE method (Spoofing, Tampering, Repudiation, Information Disclosure, Denial of Service, and Elevation of Privilege). It adds nodes to each subgraph representing the different components of the system being modeled and prefixes the name of the component with the corresponding threat type. Then it adds edges between the nodes to show the relationships between the components. Finally, it adds the subgraphs to the main graph and generates the diagram in the desired format
*/

func main() {
	// Create a new graph object
	g := gographviz.NewGraph()
	g.SetName("G")
	g.SetDir(true)
	// Threat Subgraphs
	spoofingGraph := gographviz.NewGraph()
	tamperingGraph := gographviz.NewGraph()
	repudiationGraph := gographviz.NewGraph()
	informationDisclosureGraph := gographviz.NewGraph()
	denialOfServiceGraph := gographviz.NewGraph()
	elevationOfPrivilegeGraph := gographviz.NewGraph()
	// Add nodes to the graph with the threat type as prefix
	spoofingGraph.AddNode("", "Spoofing: Node1", nil)
	tamperingGraph.AddNode("", "Tampering: Node2", nil)
	repudiationGraph.AddNode("", "Repudiation: Node3", nil)
	informationDisclosureGraph.AddNode("", "Information Disclosure: Node4", nil)
	denialOfServiceGraph.AddNode("", "Denial of Service: Node5", nil)
	elevationOfPrivilegeGraph.AddNode("", "Elevation of Privilege: Node6", nil)
	// Add edges between the nodes
	spoofingGraph.AddEdge("Spoofing: Node1", "Tampering: Node2", true, nil)
	tamperingGraph.AddEdge("Tampering: Node2", "Repudiation: Node3", true, nil)
	repudiationGraph.AddEdge("Repudiation: Node3", "Information Disclosure: Node4", true, nil)
	informationDisclosureGraph.AddEdge("Information Disclosure: Node4", "Denial of Service: Node5", true, nil)
	denialOfServiceGraph.AddEdge("Denial of Service: Node5", "Elevation of Privilege: Node6", true, nil)
	// Add subgraphs to the main graph
	g.AddSubGraph("G", spoofingGraph)
	g.AddSubGraph("G", tamperingGraph)
	g.AddSubGraph("G", repudiationGraph)
	g.AddSubGraph("G", informationDisclosureGraph)
	g.AddSubGraph("G", denialOfServiceGraph)
	g.AddSubGraph("G", elevationOfPrivilegeGraph)
	// Generate the diagram in the desired format
	s := g.String()
	gographviz.RenderString(s, "png", "out.png")
	fmt.Println("Diagram generated successfully")

}
