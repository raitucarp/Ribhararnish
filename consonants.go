package main

import (
	"math/rand"
	"time"
)

type Consonant struct {
	symbol    string
	sample    []string
	variation []string
}

func (c Consonant) getSymbol() string {
	return c.symbol
}

func (c Consonant) pickRandomSample() string {
	rand.Seed(time.Now().Unix())
	return c.sample[rand.Intn(len(c.sample))]
}

var consonants = []Consonant{
	{"b", []string{"bad", " lab"}, []string{"b", "b"}},
	{"d", []string{"did", " lady"}, []string{"d", "d", "d"}},
	{"f", []string{"find", " if"}, []string{"f", "f"}},
	{"g", []string{"give", " flag"}, []string{"g", "g"}},
	{"h", []string{"how", " hello"}, []string{"h", "h"}},
	{"j", []string{"yes", " yellow"}, []string{"y", "y"}},
	{"k", []string{"cat", " back"}, []string{"c", "ck"}},
	{"l", []string{"leg", " little"}, []string{"l", "l", "l"}},
	{"m", []string{"man", " lemon"}, []string{"m", "m"}},
	{"n", []string{"no", " ten"}, []string{"n", "n"}},
	{"ŋ", []string{"sing", " finger"}, []string{"ng", "n"}},
	{"p", []string{"pet", " map"}, []string{"p", "p"}},
	{"r", []string{"red", " try"}, []string{"r", "r"}},
	{"s", []string{"sun", " miss"}, []string{"s", "ss"}},
	{"ʃ", []string{"she", " crash"}, []string{"sh", "sh"}},
	{"t", []string{"tea", " getting"}, []string{"t", "tt"}},
	{"tʃ", []string{"check", " church"}, []string{"ch", "ch", "ch"}},
	{"θ", []string{"think", " both"}, []string{"th", "th"}},
	{"ð", []string{"this", " mother"}, []string{"th", "th"}},
	{"v", []string{"voice", " five"}, []string{"v", "ve"}},
	{"w", []string{"wet", " window"}, []string{"w", "w", "w"}},
	{"z", []string{"zoo", " lazy"}, []string{"z", "z"}},
	{"ʒ", []string{"pleasure", " vision"}, []string{"s", "si"}},
	{"dʒ", []string{"just", " large"}, []string{"j", "ge"}},
}
