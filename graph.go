package graph

import(
	"container/heap"
)

type Graph struct {
	Node map[string]Node
	Edge []Edge
}

/**
 * Node will take in a string representationa and save
 * the nodes input as an interface. There is then an array
 * of outbound edges, which which will be used to connect with other nodes
 */
type Node struct {
	Code		string
	Data 		interface{}
	OutEdges	[]Edge
}

/**
 * Structure represts edges between nodes.
 */
type Edge struct {
	Destination	*Node
	Weight		float64
	Len			int
}

/**
 * @param node - the node for the edge to be added to
 * @param edge - an edge to add
 * Adds a edge to a given node
 */
func AddOutEdge(node *Node, edge Edge){
	node.OutEdges = append(node.OutEdges, edge)
}

/**
 * Adds a node to a graph
 * @param graph - the graph to add node to
 * @param node - the node to add a graph to
 * @param code - the code to use as the map key
 */
func AddNode(g *Graph, node Node, code string){
	g.Node[code] = node
}

/*************************************************
 * The following is used for Dijkstra's algorithm
 * Implmenting a queue and D_Node which will be
 * used to replace the standard nodes. This allows
 * flexability and portation.
 *************************************************/

type D_Node struct {
	Node		Node
	Value 		int
	Back 		*D_Node
}

type Queue []*D_Node

func (q Queue) Len() int { return len(q) }

func (q Queue) Less(i, j int) bool {
	return q[i].Value > q[j].Value
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x interface{}) {
	node := x.(*D_Node)
	*q = append(*q, node)
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	node := old[n-1]
	*q = old[0 : n-1]
	return node
}

/**
 * Uses Dijkstra's Algorithm to search through the graph
 * @param g - graph for to search
 * @param source - node to search from (starting node)
 * @param destination - node to search for (end node)
 * @return list of nodes in reverse order
 */
func Dijkstra(g *Graph, source *Node, destination *Node) ([]Node, int) {

	var edge Edge
	var path []Node
	var dnodes = make(map[string]*D_Node)
	var dcurrent *D_Node

	srccode := source.Code

	for code, node := range g.Node {
		dnodes[code] = &D_Node{node, 99999999, nil}
	}

	n := &Queue{}
	heap.Init(n)

	dcurrent = dnodes[srccode]
	dcurrent.Value = 0

	heap.Push(n, dcurrent)

	for n.Len() > 0 {

		dcurrent = heap.Pop(n).(*D_Node)

		for i := 0; i < len(dcurrent.Node.OutEdges); i++ {

			edge = dcurrent.Node.OutEdges[i]

			if dnodes[edge.Destination.Code].Value > edge.Len + dcurrent.Value {

				heap.Push(n, dnodes[edge.Destination.Code])

				dnode := dnodes[edge.Destination.Code]
				dnode.Value = edge.Len + dcurrent.Value
				dnode.Back = dcurrent
			}
		}
	}

	dcurrent = dnodes[destination.Code]

	for {
		if dcurrent == nil { break }
		path = append(path, dcurrent.Node)
		if dcurrent.Node.Code == srccode { break }
		dcurrent = dcurrent.Back
	}

	return path, dnodes[destination.Code].Value
}
