package graph

import ( "testing" )


// Tests graph construction
 
func TestGraphConstruction(t *testing.T) {

  str := "TEST"
  node := Node{str, nil, nil}
  
  g := Graph{}
  g.Node = make(map[string]Node)
  g.Node[str] = node
  
  if(str != g.Node[str].Code){
          t.Errorf("Recieved %v, expected %v", g.Node[str].Code, str)
  }
}


// Tests edge creation
 
func TestEdgeCreation (t *testing.T) {

  g := Graph{}
  g.Node = make(map[string]Node)
  
  str := "TEST1"
  
  g.Node[str] = Node{str, test{1}, nil}
  
  var edgeone []Edge
  nodeone := g.Node[str]
  edgeone = append(edgeone, Edge{&nodeone, 0, 10})
  
  if(edgeone[0].Destination.Code != str){
          t.Errorf("Recieved %v, expected %v",edgeone[0].Destination.Code, str)
  }
}


// Tests adding an edge

func TestAddEdge (t *testing.T) {

  g := Graph{}
  g.Node = make(map[string]Node)
  
  g.Node["TEST1"] = Node{"TEST1", test{1}, nil}
  
  nodeone := g.Node["TEST1"]
  g.Node["TEST2"] = Node{"TEST2", test{2}, nil}
  node := g.Node["TEST2"]
  AddOutEdge(&node, Edge{&nodeone, 0, 10})
  
  if(node.OutEdges[0].Destination.Code != "TEST1"){
          t.Errorf("Recieved %v, expected %v",g.Node["TEST2"].OutEdges[0].Destination.Code, "TEST1")
  }
}


// Tests adding a node

func TestAddNode (t *testing.T) {
  
  g := Graph{}
  g.Node = make(map[string]Node)
  AddNode(&g, Node{"TEST1", test{1}, nil}, "TEST1")
  
  if(g.Node["TEST1"].Code != "TEST1"){
          t.Errorf("Recieved %v, expected %v\n", g.Node["TEST1"].Code, "TEST1")
  }
}


// Test dijkstra's algorithm

type test struct {
        Value int
}
func TestDijkstra (t *testing.T){

  g := Graph{}
  g.Node = make(map[string]Node)
  
  g.Node["TEST1"] = Node{"TEST1", test{1}, nil}
  
  var edgeone []Edge
  nodeone := g.Node["TEST1"]
  edgeone = append(edgeone, Edge{&nodeone, 0, 10})
  g.Node["TEST2"] = Node{"TEST2", test{2}, edgeone}
  
  var edgetwo []Edge
  nodeone = g.Node["TEST1"]
  edgetwo = append(edgetwo, Edge{&nodeone, 0, 20})
  g.Node["TEST3"] = Node{"TEST3", test{3}, edgetwo}
  
  var edgethree []Edge
  nodeone = g.Node["TEST2"]
  nodetwo := g.Node["TEST3"]
  edgethree = append(edgethree, Edge{&nodeone, 0, 1})
  edgethree = append(edgethree, Edge{&nodetwo, 0, 1})
  g.Node["TEST4"] = Node{"TEST4", test{4}, edgethree}
  
  
  src := g.Node["TEST4"]
  dest := g.Node["TEST1"]
  
  var path []test
  
  path = append(path, test{1})
  path = append(path, test{2})
  path = append(path, test{4})
  
  check, _ := Dijkstra(&g, &src, &dest)
  
  for j := 0; j < len(check); j++ {
    if(path[j].Value != check[j].Data.(test).Value){
      t.Errorf("Recieved %v, expected %v", path[j].Value, check[j].Code)
    }
  }
}
