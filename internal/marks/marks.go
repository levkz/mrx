package marks

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
