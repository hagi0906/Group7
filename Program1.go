package main

import (
	//"fmt"
	//	"strconv"
	"strings"
)

func inttoroman(i int) string {

	var roman string
	alabtable := []int{4000, 3000, 2000, 1000, 900, 800, 700, 600, 500, 400, 300, 200, 100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	romantable := []string{"MMMM", "MMM", "MM", "M", "CM", "DCCC", "DCC", "DC", "D", "CD", "CCC", "CC", "C", "XC", "LXXX", "LXX", "LX", "L", "XL", "XXX", "XX", "X", "IX", "VIII", "VII", "VI", "V", "IV", "III", "II", "I"}

	tablesize := 31
	var romanresult []string

	for j := 0; j < tablesize; j++ {
		result := i / alabtable[j]

		if result == 1 {

			romanresult = append(romanresult, romantable[j])
			i = i - alabtable[j]
		}
	}
	roman = strings.Join(romanresult, "")

	return roman
}

/*
func main() {
	x := 1990

	fmt.Println(inttoroman(x))

}
*/
