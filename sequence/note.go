package sequence

import "fmt"

type Note struct {
	NameSharp string
	NameFlat  string
	Next      *Note
	Previous  *Note
}

func (n *Note) FindFromSemitones(semitonesDistance int) *Note {
	curr := n
	for range semitonesDistance {
		curr = curr.Next
	}
	return curr
}

func (n *Note) GetIntervalInSemitones(note *Note) int {
	interval := 0

	for {
		if note.NameFlat == n.NameFlat {
			return interval
		}
		interval++
	}
}

func (n *Note) PrintSharpToTerminal() {
	fmt.Printf("%s\n", n.NameSharp)
}

func (n *Note) PrintFlatToTerminal() {
	fmt.Printf("%s\n", n.NameFlat)
}
