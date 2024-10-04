package trie

type Node struct {
	value string
	nodes map[string]*Node
	match bool
}

func NewTrie() *Node {
	return &Node{
		value: "",
		nodes: make(map[string]*Node, 10),
		match: false,
	}
}

func (node *Node) AddValue(value string) *Node {
	currentNode := node

	for i := 0; i < len(value); i++ {
		char := value[i : i+1]
		charNode, nodeExists := currentNode.nodes[char]

		if !nodeExists {
			charNode = NewTrie()
			currentNode.nodes[char] = charNode
		}

		currentNode = charNode
	}

	if currentNode != node {
		currentNode.value = value
		currentNode.match = true
	}

	return node
}

func (node *Node) AddValuesFromArray(values []string) *Node {
	for i := 0; i < len(values); i++ {
		node.AddValue(values[i])
	}

	return node
}

func (node *Node) Search(value string) *Node {
	currentNode := node

	for i := 0; i < len(value); i++ {
		char := value[i : i+1]
		charNode, nodeExists := currentNode.nodes[char]

		if !nodeExists {
			currentNode = nil
			break
		}

		currentNode = charNode
	}

	return currentNode
}

func (node *Node) Matched() bool {
	return node.match
}

func (node *Node) Size() int {
	return len(node.nodes)
}
