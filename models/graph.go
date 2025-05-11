package models

import "math"

const EPSILON = 1e-6

type Edge struct {
	To     int
	Rev    int
	Cap    float64
	Origin int
}

type Graph struct {
	N     int
	Adj   [][]Edge
	Level []int
	Iter  []int
	Names map[int]string
	Ids   map[string]int
}

func NewGraph(size int) *Graph {
	return &Graph{
		N:     size,
		Adj:   make([][]Edge, size),
		Level: make([]int, size),
		Iter:  make([]int, size),
		Names: make(map[int]string),
		Ids:   make(map[string]int),
	}
}

func (g *Graph) AddEdge(from, To int, Cap float64) {
	g.Adj[from] = append(g.Adj[from], Edge{To, len(g.Adj[To]), Cap, from})
	g.Adj[To] = append(g.Adj[To], Edge{from, len(g.Adj[from]) - 1, 0, To})
}

func (g *Graph) BFS(s, t int) {
	for i := range g.Level {
		g.Level[i] = -1
	}
	queue := []int{s}
	g.Level[s] = 0

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for _, e := range g.Adj[v] {
			if e.Cap > EPSILON && g.Level[e.To] < 0 {
				g.Level[e.To] = g.Level[v] + 1
				queue = append(queue, e.To)
			}
		}
	}
}

func (g *Graph) DFS(v, t int, upTo float64) float64 {
	if v == t {
		return upTo
	}
	for i := g.Iter[v]; i < len(g.Adj[v]); g.Iter[v], i = g.Iter[v]+1, i+1 {
		e := &g.Adj[v][i]
		if e.Cap > EPSILON && g.Level[v] < g.Level[e.To] {
			d := g.DFS(e.To, t, math.Min(upTo, e.Cap))
			if d > EPSILON {
				e.Cap -= d
				g.Adj[e.To][e.Rev].Cap += d
				return d
			}
		}
	}
	return 0
}

func (g *Graph) MaxFlow(s, t int) float64 {
	flow := 0.0
	for {
		g.BFS(s, t)
		if g.Level[t] < 0 {
			break
		}
		for i := range g.Iter {
			g.Iter[i] = 0
		}
		for {
			f := g.DFS(s, t, math.MaxFloat64)
			if f < EPSILON {
				break
			}
			flow += f
		}
	}
	return flow
}
