package t5tokenizer

type Token struct {
	Surface  string
	Features []string
}

func (t *Token) Append(s string) {
	t.Surface = t.Surface + s
}

func (t *Token) Empy() {
	t.Surface = ""
}

func (t *Token) AppendFeature(f string) {
	t.Features = append(t.Features, f)
}
