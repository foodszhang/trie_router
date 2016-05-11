package router

// Node struct define the tree
type Node struct {
	childNodes map[rune]*Node
	Char       rune
	Exist      bool
}

//CreateNode init a new  node
func CreateNode() *Node {
	node := new(Node)
	node.childNodes = make(map[rune]*Node)
	return node
}

// Insert insert a string into the tree
func (root *Node) Insert(s string) {
	node := root
	for _, r := range s {
		if node.childNodes[r] == nil {
			node.childNodes[r] = CreateNode()
		}
		node = node.childNodes[r]
	}
	node.Exist = true
}

// Search for the string if in the tree
func (root *Node) Search(s string) bool {
	node := root
	for _, r := range s {
		node = node.childNodes[r]
		if node == nil {
			return false
		}
	}
	if node.Exist {
		return true
	} else {
		return false
	}
}
