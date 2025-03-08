package graphs

import "fmt"


//Writing an industrial strength undirected graph implementation:
// 1. Can construct the graph with v vertices with no edges
// 2. Can read a stream of vertices in the form of key value pair representing an edge and create a graph
// 3. Can add an edge to a graph 
// 4. Iterate over the vertices adjacent to a given vertex

//Implementing an undirected graph
// 1. Will use an adjacency list representation
// 2. Will use a slice of slices to represent the adjacency list
// 3. The outer slice will have the number of vertices in the graph
// 4. The inner slice will have the vertices adjacent to the vertex



type UndirectedGraph struct {
	v int
	e int
	adj [][]int
}

func (g *UndirectedGraph) AddEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.e++
}

func (g *UndirectedGraph) Adj(v int) []int {
	return g.adj[v]
}

func (g *UndirectedGraph) V() int {
	return g.v
}

func (g *UndirectedGraph) E() int {
	return g.e
}

func (g *UndirectedGraph) String() string {
	return fmt.Sprintf("Vertices: %d, Edges: %d", g.v, g.e)
}

func NewUndirectedGraph(v int) *UndirectedGraph {
	return &UndirectedGraph{
		v: v,
		e: 0,
		adj: make([][]int, v),
	}
}

func NewUndirectedGraphFromStream(v int, stream [][]int) *UndirectedGraph {
	g := NewUndirectedGraph(v)
	for _, edge := range stream {
		g.AddEdge(edge[0], edge[1])
		g.AddEdge(edge[1], edge[0])
	}
	return g
}