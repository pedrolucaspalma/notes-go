package main

import "github.com/pedrolucaspalma/notes-go/sequence"

func main() {
	sequence := sequence.NewNoteSequence()

	notes, _ := sequence.PentatonicScale("C")

	for _, n := range notes {
		n.PrintFlatToTerminal()
	}
}
