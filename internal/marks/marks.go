package marks

import "fmt"

type Mark struct {
	Label string
	Path  string
}

func NewMark(label string, path string) Mark {
	return Mark{
		Path:  path,
		Label: label,
	}
}

func (m *Mark) Print() {
	fmt.Print("Label: ", m.Label, "\tPath:", m.Path, "\n")
}
