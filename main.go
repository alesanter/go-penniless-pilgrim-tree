package main

import (
	"fmt"
	"os"
)

func main() {
	const SIZE = 5
	InitNode(SIZE)
	InitBlock(SIZE)

	p := &Path{
		SIZE,
		[]*Block{
			GetBlock(GetNode(0, 0), DIR_RIGHT),
			GetBlock(GetNode(0, 1), DIR_RIGHT),
		},
	}
	ps := []*Path{p}

	for {
		count := len(ps)
		if count == 0 {
			break
		}
		fmt.Fprintln(os.Stderr, count)

		ps1 := make([]*Path, 0, len(ps))
		for _, p := range ps {
			p1s := p.Evolve()

			for _, p1 := range p1s {
				n := p1.LastNode()

				if n == GetNode(SIZE-1, SIZE-1) {
					tax := p1.Tax()
					if tax < 1e-10 {
						fmt.Println(tax, p1.Dirs())
					}
				} else {
					ps1 = append(ps1, p1)
				}
			}
		}

		ps = ps1
	}
}
