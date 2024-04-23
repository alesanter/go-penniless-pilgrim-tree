package main

import "fmt"

type Block struct {
	Head *Node
	Tail *Node
	Dir  Dir
}

func (b *Block) String() string {
	return fmt.Sprintf("{%v %v}", b.Head, b.Dir)
}

var AllBlocks map[*Node]map[Dir]*Block

func InitBlock(s int) {
	AllBlocks = make(map[*Node]map[Dir]*Block, s*s)
	for r, ns := range AllNodes {
		for c, h := range ns {
			AllBlocks[h] = make(map[Dir]*Block)

			ds := map[Dir]bool{DIR_RIGHT: true, DIR_LEFT: true, DIR_UP: true, DIR_DOWN: true}
			if r == 0 {
				delete(ds, DIR_UP)
			} else if r == s-1 {
				delete(ds, DIR_DOWN)
			}
			if c == 0 {
				delete(ds, DIR_LEFT)
			} else if c == s-1 {
				delete(ds, DIR_RIGHT)
			}

			for d := range ds {
				t := h.Move(d)
				AllBlocks[h][d] = &Block{h, t, d}
			}
		}
	}
}

// Nullable
func GetBlock(n *Node, d Dir) *Block {
	bs, ok := AllBlocks[n]
	if !ok {
		return nil
	}
	b, ok := bs[d]
	if !ok {
		return nil
	}
	return b
}

func (b *Block) Equal(b1 *Block) bool {
	return b.Head == b1.Head && b.Dir == b1.Dir || b.Tail == b1.Head && b.Dir == b1.Dir.Opposite()
}
