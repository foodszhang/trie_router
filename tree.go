package trie_router

type Node struct {
	childNodes map[rune]*Node
	Char       rune
	Exist      bool
}

func CreateNode() *Node {
	node := new(Node)
	node.childNodes = make(map[rune]*Node)
	return node
}
func (root *Node) Insert(s *string) {
	node := root
	for _, r := range *s {
		if node.childNodes[r] == nil {
			node.childNodes[r] = CreateNode()
		}
		node = node.childNodes[r]
	}
	node.Exist = true
}
