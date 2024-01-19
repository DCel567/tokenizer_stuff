package main

import (
	//t5 "test/t5tokenizer"

	"fmt"
	"log"
	"time"

	"github.com/sugarme/tokenizer"
	"github.com/sugarme/tokenizer/model/bpe"
	"github.com/sugarme/tokenizer/pretokenizer"
)

func train() {
	startTime := time.Now()

	files := []string{
		// "example/tokenizer/bpe/train/input/oscar.eo-50k.txt",
		// "example/tokenizer/bpe/train/input/adieu.txt",
		// "example/tokenizer/bpe/train/input/test.txt",
		// "example/tokenizer/bpe/train/input/test-eo.txt",

		// "input/oscar.eo.txt",
		// "input/epo_literature_2011_300K-sentences.txt",
		// "input/epo_mixed_2012_1M-sentences.txt",
		// "input/epo_newscrawl_2017_1M-sentences.txt",
		// "input/epo_web_2011_100K-sentences.txt",
		// "input/epo_web_2012_1M-sentences.txt",
		// "input/epo_wikipedia_2007_300K-sentences.txt",
		// "input/epo_wikipedia_2011_300K-sentences.txt",
		// "input/epo_wikipedia_2012_300K-sentences.txt",
		// "input/epo_wikipedia_2016_300K-sentences.txt",

		"corpus.txt",
	}

	var vocab map[string]int = make(map[string]int)
	vocab["<s>"] = 0
	vocab["<pad>"] = 1
	vocab["</s>"] = 2
	vocab["<unk>"] = 3
	vocab["<mask>"] = 4
	vocab["ˈ"] = 5
	vocab["\u0361"] = 6
	vocab["["] = 7
	vocab["ɭ"] = 8
	vocab["ɑ"] = 9
	vocab["ʲ"] = 10
	vocab["ʃ"] = 11
	vocab["ʌ"] = 12

	var merges bpe.Merges = make(map[bpe.Pair]bpe.PairVal)

	model := bpe.NewBPE(vocab, merges)

	fmt.Println(merges)

	unkToken := "<unk>"
	model.UnkToken = &unkToken

	trainer := bpe.NewBpeTrainer(3, 52000)

	tk := tokenizer.NewTokenizer(model)

	specialToks := []tokenizer.AddedToken{
		tokenizer.NewAddedToken("ˈj͡a", true),
	}
	tk.AddSpecialTokens(specialToks)

	// charSet := make(map[string]struct{})

	// charSet["ˈ"] = struct{}{}
	// charSet["\u0361"] = struct{}{}
	// charSet["ʒ"] = struct{}{}
	// charSet["ɭ"] = struct{}{}
	// charSet["ɑ"] = struct{}{}
	// charSet["ʲ"] = struct{}{}
	// charSet["ʃ"] = struct{}{}
	// charSet["ʌ"] = struct{}{}

	// fmt.Println("ʌ"[1])
	// fmt.Println(len("b"))

	// trainer.InitialAlphabet = charSet
	//fmt.Println(trainer.InitialAlphabet)

	bytelevel := pretokenizer.NewByteLevel()

	tk.WithPreTokenizer(bytelevel)

	err := tk.Train(trainer, files)
	if err != nil {
		log.Fatal(err)
	}

	trainedModel := tk.GetModel()

	trainedModel.Save("./model", "ipa")

	trainedTime := time.Since(startTime).Seconds() / 60

	fmt.Printf("Training time (min): %f.2\n", trainedTime)
}

func test() {
	model, err := bpe.NewBpeFromFiles("./model/ipa-vocab.json", "model/ipa-merges.txt")
	if err != nil {
		log.Fatal(err)
	}

	tk := tokenizer.NewTokenizer(model)

	bl := pretokenizer.NewBertPreTokenizer()

	tk.WithPreTokenizer(bl)

	//TODO: replace with reading from test.txt file
	sentence := "ˈj͡a ʒdaɭˈɑ tʲibʲˈɑ tˈɑk ʒdaɭˈɑ tˈy bˈyɭ ʒˈopa mʲit͡ʃʲtˈoj͡u mʌjˈej xrustˈɑɭnʌj͡u"

	inputSeq := tokenizer.NewInputSequence(sentence)

	en, err := tk.Encode(tokenizer.NewSingleEncodeInput(inputSeq), false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sentence: '%v'\n", sentence)

	fmt.Printf("Tokens: %+v\n", en.GetTokens())

	for _, tok := range en.GetTokens() {
		fmt.Printf("'%v'\n", tok)
	}
}

func main() {
	train()
	test()
}
