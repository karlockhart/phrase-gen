package methods

type node struct {
	prev      *node
	generator func() string
	after     func(string) string
	children  []*node
}

type Graph struct {
	root *node
}

func (g *Graph) GetRoot() *node {
	return g.root
}

func (n *node) AddNode(g func() string, a func(string) string) (c *node) {
	c = new(node)
	c.generator = g
	n.children = append(n.children, c)
	return
}

func NewGenerator() (g *Graph) {
	g = new(Graph)
	g.root = new(node)
	g.root.prev = nil
	return
}
