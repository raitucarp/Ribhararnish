package main

import (
	"math/rand"
	"time"
)

type Vowel struct {
	symbol    string
	sample    []string
	variation []string
}

func (v Vowel) getSymbol() string {
	return v.symbol
}

func (v Vowel) pickRandomSample() string {
	rand.Seed(time.Now().Unix())
	return v.sample[rand.Intn(len(v.sample))]
}

var vowels = []Vowel{
	{"ʌ", []string{"cup", " luck"}, []string{"u", "u"}},
	{"ɑ:", []string{"arm", " father "}, []string{"a", "a"}},
	{"æ", []string{"cat", " black"}, []string{"a", "a"}},
	{"e", []string{"met", " bed"}, []string{"e", "e"}},
	{"ə", []string{"away", " cinema"}, []string{"a", "e", "a"}},
	{"ɜ:ʳ", []string{"turn", " learn"}, []string{"ur", "ear"}},
	{"ɪ", []string{"hit", " sitting"}, []string{"i", "i", "i"}},
	{"i:", []string{"see", " heat"}, []string{"ee", "ea"}},
	{"ɒ", []string{"hot", " rock"}, []string{"o", "o"}},
	{"ɔ:", []string{"call", " four"}, []string{"a", "ou"}},
	{"ʊ", []string{"put", " could"}, []string{"u", "oul"}},
	{"u:", []string{"blue", " food"}, []string{"ue", "oo"}},
	{"aɪ", []string{"five", " eye"}, []string{"i", "eye"}},
	{"aʊ", []string{"now", " out"}, []string{"ow", "ou"}},
	{"eɪ", []string{"say", " eight"}, []string{"ay", "eigh"}},
	{"oʊ", []string{"go", " home"}, []string{"o", "o"}},
	{"ɔɪ", []string{"boy", " join"}, []string{"oy", "oi"}},
	{"eəʳ", []string{"where", " air"}, []string{"ere", "air"}},
	{"ɪəʳ", []string{"near", " here"}, []string{"ear", "ere"}},
	{"ʊəʳ", []string{"pure", " tourist"}, []string{"ure", "our"}},
}
