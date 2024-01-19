package t5tokenizer

import (
	"bufio"
	"log"
	"os"
)

type BPELearner struct {
	Symbols      int
	MinFrequency int
	DictInput    bool // true if dictionary present. implicit loadFromDictionary call
	TotalSymobls bool
	Vocab        map[string]int
}

func NewBPELearner(s, mf int, di, ts bool) *BPELearner {
	l := new(BPELearner)
	l.Symbols = s
	l.MinFrequency = mf
	l.DictInput = di
	l.TotalSymobls = ts
	l.Vocab = make(map[string]int)
	return l
}

func (l *BPELearner) LoadFromDictionary(f string) {
	//  f is vocabulary file from which Vocab generated
	//  vocabulary file structure: "token number\n"
	//  what number stands for isn't clear
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	// for sc.Scan() {
	// 	str := strings.Split(sc.Text(), " ")
	// 	tk := str[0]
	// 	fq, _ := strconv.Atoi(str[1])

	// 	(l.Vocab)[tk] = fq
	// }

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
}

//func (l *BPELearner)
