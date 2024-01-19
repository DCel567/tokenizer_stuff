package t5tokenizer

import (
	"bufio"
	"log"
	"os"
)

type Tokenizer struct {
	Joiner_marker            string
	Spacer_marker            string
	Ph_marker_open           string
	Ph_marker_close          string
	Escaped_character_prefix string
	Escaped_character_width  int8
}

func NewTokenizer() *Tokenizer {
	t := new(Tokenizer)
	t.Joiner_marker = "￭"
	t.Spacer_marker = "▁"
	t.Ph_marker_open = "｟"
	t.Ph_marker_close = "｠"
	t.Escaped_character_prefix = "％"
	t.Escaped_character_width = 4
	return t
}

func (t *Tokenizer) Tokenize(f string) {
	// f for source file
	// takes text in file and returns
	// array of tokens

	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		//sc.Text() inside

		for i, ch := range sc.Text() {

		}
	}
}

func (t *Tokenizer) Detokenize(tokens []Token) string {
	return "jopa"
}
