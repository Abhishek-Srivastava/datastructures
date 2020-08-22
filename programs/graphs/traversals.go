package graphs

import (
	"datastructures/programs/queues"
	"datastructures/programs/stack"
	"fmt"
	"math"
)

// Graph represents a graph with Adjancency list
type Graph struct {
	Adjancenylist map[int][]map[int]int
}

// NewGraph creates a new instance
func NewGraph() *Graph {
	// adjList := make(map[int][]map[int]int)
	return &Graph{Adjancenylist: nil}
}

// AddVertex se
func (g *Graph) AddVertex(v int) {
	_, ok := g.Adjancenylist[v]
	if !ok {
		if g.Adjancenylist == nil {
			g.Adjancenylist = make(map[int][]map[int]int)
			g.Adjancenylist[v] = nil
		}
		g.Adjancenylist[v] = append(g.Adjancenylist[v], nil)
		if len(g.Adjancenylist[v][0]) == 0 {
			g.Adjancenylist[v] = g.Adjancenylist[v][1:len(g.Adjancenylist[v])]
		}
	}
}

func contains(l []map[int]int, v int) (bool, int) {
	for index, item := range l {
		_, ok := item[v]
		if ok {
			return true, index
		}
	}
	return false, -1
}

// AddEdge se
func (g *Graph) AddEdge(v1, v2, w int) {
	l, ok := g.Adjancenylist[v1]
	if ok {
		if ok, _ := contains(l, v2); !ok {
			edgeInfo := make(map[int]int)
			edgeInfo[v2] = w
			l = append(l, edgeInfo)
			g.Adjancenylist[v1] = l
		}
	}
}

// BFS se
func (g *Graph) BFS(startingNode int) {
	q := &queues.Q{
		Items: []int{},
	}
	visited := make(map[int]bool)
	q.Insert(startingNode)
	visited[startingNode] = true

	for {
		if len(q.Items) == 0 {
			break
		}
		curNode, err := q.Delete()
		if err != nil {
			fmt.Println("Unexpected error occurred")
			panic(err)
		}
		fmt.Printf("%d ", curNode)
		neighbours := g.Adjancenylist[curNode]
		for _, n := range neighbours {
			for k := range n {
				_, ok := visited[k]
				if !ok {
					q.Insert(k)
					visited[k] = true
				}
			}
		}
	}
	fmt.Println(" - BFS traversal of the graph")
}

// DFS se
func (g *Graph) DFS(startingNode int) {
	s := stack.New()
	s.Push(startingNode)
	visited := make(map[int]bool)

	for {
		if len(s.Val) == 0 {
			break
		}
		item, err := s.Pop()
		if err != nil {
			fmt.Println("Unexpected error occurred")
			panic(err)
		}
		fmt.Printf("%d ", item)
		visited[item] = true
		neighbours := g.Adjancenylist[item]
		for _, n := range neighbours {
			for k := range n {
				_, ok := visited[k]
				if !ok {
					s.Push(k)
					visited[k] = true
				}
			}
		}
	}
	fmt.Println(" - DFS traversal of the graph")
}

// CreateEdges se
func (g *Graph) CreateEdges(edgeInput [][]int) {
	for i := 0; i < len(edgeInput); i++ {
		g.AddEdge(edgeInput[i][0], edgeInput[i][1], edgeInput[i][2])
	}
}

func (g *Graph) relax(distance map[int]int, u int, visited []int) map[int]int {
	fmt.Println("Relaxing neighbours of vertex", u)
	fmt.Println("Distance matrix", distance)
	for _, neighbourMap := range g.Adjancenylist[u] {
		for v, c := range neighbourMap {
			if containsArray(visited, v) {
				continue
			}
			if distance[u]+c < distance[v] {
				distance[v] = distance[u] + c
			}
		}
	}
	return distance
}

// containsArray
func containsArray(a []int, val int) bool {
	for _, items := range a {
		if items == val {
			return true
		}
	}
	return false
}

func selectVertexWithMinVal(distance map[int]int, exludeItems []int) (int, int) {
	min := math.MaxInt32
	index := -1
	for key, val := range distance {
		if containsArray(exludeItems, key) {
			continue
		}
		if val < min {
			min = val
			index = key
		}
	}
	return min, index
}

func (g *Graph) dijstrasShortestPath(startingNode int, distance map[int]int) {
	visited := []int{}
	// initialize the distance matrix
	for _, neighbours := range g.Adjancenylist[nodes[0]] {
		fmt.Println("looking neighbour:", neighbours)
		for node, cost := range neighbours {
			distance[node] = cost
		}
	}
	// distance[2] = 50
	// distance[3] = 45
	// distance[4] = 10
	// distance[5] = math.MaxInt32
	// distance[6] = math.MaxInt32
	for len(visited) < len(nodes) {
		_, visitedItem := selectVertexWithMinVal(distance, visited)
		if visitedItem == -1 {
			break
		}
		fmt.Println("visited item", visitedItem)
		distance = g.relax(distance, visitedItem, visited)
		visited = append(visited, visitedItem)
		fmt.Println("Distance matrix after relaxation of :", visitedItem, distance)
	}
	fmt.Println("Distances from vertex ", startingNode, distance)
}

var nodes = []int{1, 2, 3, 4, 5, 6}

// Run se
func Run() {
	g := NewGraph()
	for _, i := range nodes {
		g.AddVertex(i)
	}
	// edges := [][]int{
	// 	{1, 2, 50},
	// 	{1, 3, 45},
	// 	{1, 4, 10},
	// 	{2, 3, 10},
	// 	{2, 4, 15},
	// 	{3, 5, 30},
	// 	{4, 1, 10},
	// 	{4, 5, 15},
	// 	{5, 2, 20},
	// 	{5, 3, 35},
	// 	{6, 5, 3},
	// }
	edges := [][]int{
		{1, 2, 2},
		{1, 3, 4},
		{2, 3, 1},
		{2, 4, 7},
		{3, 5, 3},
		{4, 6, 1},
		{5, 4, 2},
		{5, 6, 5},
	}
	g.CreateEdges(edges)
	fmt.Println("Adj list")
	for k, v := range g.Adjancenylist {
		fmt.Println(k, v)
	}
	g.BFS(nodes[0])
	g.DFS(nodes[0])
	distance := make(map[int]int)
	i := 1
	for i < len(nodes) {
		distance[nodes[i]] = math.MaxInt32
		i++
	}
	g.dijstrasShortestPath(nodes[0], distance)
}
