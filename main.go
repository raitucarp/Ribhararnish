package main

import "fmt"

func convertStructureToList(s []string) [][]Phonetic {
	var gaga [][]Phonetic
	for _, x := range s {
		switch x {
		case "v":
			var g []Phonetic
			for _, v := range vowels {
				g = append(g, v)
			}
			gaga = append(gaga, g)
			break
		case "c":
			var g []Phonetic
			for _, c := range consonants {
				g = append(g, c)
			}
			gaga = append(gaga, g)
			break
		}
	}
	return gaga
}

func bb(s [][]Phonetic) {

}

func main() {
	a := makeMultiDimmensionStructures()
	for _, y := range a {
		fmt.Println("------", y)
		gaga := convertStructureToList(y)

		// ss := [][]Phonetic{}
		for _, s := range gaga {
			fmt.Println("####")
			baba := []Phonetic{}

			for _, g := range s {
				baba = append(baba, g)
				// fmt.Println("g", g)
				// i := 0
				// for ok := true; ok; ok = condition {

				// }
			}

			fmt.Println("baba", baba)
		}
	}
}
