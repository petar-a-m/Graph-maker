package graph

import (
	"fmt"
	"os"
	"encoding/csv"
)

type Graph struct {
	Nodes []node
	Links []link
}
type node struct {
	id string
}

type link struct{
	node1 node
	node2 node
}

func findNodeById(x string) int {
	for i, n := range g.Nodes {
		if x == n.id {
			return i
		}
	}
	return 0
}
func createLink(n1, n2 string) {
	
	var l link
	l.node1 = g.Nodes[findNodeById(n1)]
	l.node2 = g.Nodes[findNodeById(n2)]
	g.Links = append(g.Links,l)
}
func createNode(id string) {
	var n node
	n.id = id
	g.Nodes = append(g.Nodes,n)
}
func createArrayFromCSV(fpath string) [][]string{
	file, err := os.Open(fpath)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, _ := reader.ReadAll()
	return records
}
func resetGraph(){
	g.Nodes = nil 
	g.Links = nil
}
func createGraphFromArray(a [][]string){
	for i := 0; i < len(a); i++ {
		if !graphContainsNode(a[i][0]) {
			createNode(a[i][0])
		}
		for j:= 1; j < len(a[i]); j++ {
			if !graphContainsNode(a[i][j]) {
				createNode(a[i][j])
			}
			if !graphContainsLink(a[i][0],a[i][j]) {
				createLink(a[i][0],a[i][j])
			}
		}
	}	
}
func graphContainsNode(x string) bool {
	for _, n := range g.Nodes {
		if x == n.id {
			return true
		}
	}
	return false
}
func graphContainsLink(n1, n2 string) bool {
	for _, l := range g.Links {
		if (l.node1.id == n1)&&(l.node2.id == n2) {
			return true
		}
		if (l.node2.id == n1)&&(l.node1.id == n2) {
			return true
		}
	}
	return false
}
func findAdjacentNodes(n string) []string{
	var a []string
	for _, x := range g.Nodes {
		if graphContainsLink(x.id,n){
			a = append(a,x.id)
		}
	}
	return a
}

func LoadGraphFromFile(fpath string) Graph {
	a := createArrayFromCSV(fpath)
	createGraphFromArray(a)
	return g
}

var g Graph

func main() {
	LoadGraphFromFile(os.Args[1])
	fmt.Println("Links: ", g.Links)
	fmt.Println(findAdjacentNodes("76"))
}
