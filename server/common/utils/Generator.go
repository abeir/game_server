package utils

type Generator struct {
	i int64
}

func NewGenerator() Generator {
	return Generator{0}
}

func (g *Generator) Next() int64 {
	g.i++
	return g.i
}
