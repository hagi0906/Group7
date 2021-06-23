package main

import (
	//"fmt"
	"strings"
)

func RomantoArab(s string) int {
	alabtable := []int{4000, 3000, 2000, 1000, 900, 800, 700, 600, 500, 400, 300, 200, 100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	romantable := []string{"MMMM", "MMM", "MM", "M", "CM", "DCCC", "DCC", "DC", "D", "CD", "CCC", "CC", "C", "XC", "LXXX", "LXX", "LX", "L", "XL", "XXX", "XX", "X", "IX", "VIII", "VII", "VI", "V", "IV", "III", "II", "I"}

	tablesize := 31
	arab := 0

	for i := 0; i < tablesize; i++ { //対応表でループ
		if strings.Index(s, romantable[i]) == 0 {
			arab = arab + alabtable[i]
			s = s[len(romantable[i]):len(s)]
		}

	}

	return arab
}

/*
func main() {
	x := "MMMMCCI"
	fmt.Println(RomantoArab(x))

}
*/
