package main

var stringStructures = []string{
	"v",
	"vc", "cv",
	"vcv", "cvc", "vvc", "ccv", "cvv",
	"vcvc", "cvcv",
}

func makeMultiDimmensionStructures() (m [][]string) {
	for _, s := range stringStructures {
		var y []string
		for _, a := range s {
			y = append(y, string(a))
		}
		m = append(m, y)
	}
	return m
}

type Phonetic interface {
	getSymbol() string
	pickRandomSample() string
}
