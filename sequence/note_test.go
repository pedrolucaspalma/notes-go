package sequence

import "testing"

func TestFindFromSemitones(t *testing.T) {
	sequence := NewNoteSequence()

	type table struct {
		initial  string
		interval int
		expected string
	}

	cases := []table{
		{
			initial:  "C",
			interval: 0,
			expected: "C",
		},
		{
			initial:  "C",
			interval: 1,
			expected: "C#",
		},
		{
			initial:  "C",
			interval: 5,
			expected: "F",
		},
		{
			initial:  "C",
			interval: 12,
			expected: "C",
		},
		{
			initial:  "C",
			interval: 13,
			expected: "C#",
		},
	}

	for i, c := range cases {
		initialNote := sequence.first
		if c.initial != initialNote.NameSharp {
			t.Fatalf("initial note incorrect. Case number %d. Expected %s. Got %s", i, c.initial, initialNote.NameFlat)
		}

		targetNote := initialNote.FindFromSemitones(c.interval)

		if targetNote.NameSharp != c.expected {
			t.Fatalf("got incorrect note. Case number %d. Expected %s. Got %s", i, c.expected, targetNote.NameSharp)
		}
	}
}
