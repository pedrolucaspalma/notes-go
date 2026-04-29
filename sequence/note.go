package sequence

import "fmt"

type Note struct {
	NameSharp string
	NameFlat  string
	Next      *Note
	Previous  *Note
}

func (n *Note) PrintSharpToTerminal() {
	fmt.Printf("%s\n", n.NameSharp)
}

func (n *Note) PrintFlatToTerminal() {
	fmt.Printf("%s\n", n.NameFlat)
}
