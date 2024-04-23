package main

import "fmt"

type Path struct {
	Size   int
	Blocks []*Block
}

func (p *Path) String() string {
	return fmt.Sprintf("{%d %v}", p.Size, p.Blocks)
}

func (p *Path) Contains(b *Block) bool {
	for _, b1 := range p.Blocks {
		if b.Equal(b1) {
			return true
		}
	}
	return false
}

// Panicable
func (p *Path) LastBlock() *Block {
	return p.Blocks[len(p.Blocks)-1]
}

// Panicable <- Path.LastBlock
func (p *Path) LastNode() *Node {
	b := p.LastBlock()
	return b.Tail
}

// Panicable <- Path.LastBlock
func (p *Path) Evolve() []*Path {
	b := p.LastBlock()
	n := b.Tail
	if n == nil {
		return nil
	}

	ps := make([]*Path, 0, 4)
	od := b.Dir.Opposite()
	for _, d := range AllDirs {
		if d == od {
			continue
		}

		n1 := n.Move(d)
		if n1 == nil {
			continue
		}
		b1 := GetBlock(n, d)
		if b1 == nil {
			continue
		}

		if p.Contains(b1) {
			continue
		}

		p1 := Path{p.Size, make([]*Block, 0, len(p.Blocks)+1)}
		p1.Blocks = append(p1.Blocks, p.Blocks...)
		p1.Blocks = append(p1.Blocks, b1)
		ps = append(ps, &p1)
	}

	return ps
}

func (p *Path) Tax() float64 {
	tax := 0.0
	for _, b := range p.Blocks {
		tax = b.Dir.UpdateTax(tax)
	}
	return tax
}

func (p *Path) Dirs() []Dir {
	ds := make([]Dir, len(p.Blocks))
	for i, b := range p.Blocks {
		ds[i] = b.Dir
	}
	return ds
}
