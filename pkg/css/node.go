package css


type Node struct {
	Start *Position
	End *Position
	children []*Node
	parent *Node
}
