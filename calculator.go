package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calcul(a, c int, str string) (x int) {

	switch str {
	case "+":
		x = (a + c)
	case "-":
		x = (a - c)
	case "*":
		x = (a * c)
	case "/":
		x = (a / c)
	default:
		x = 0
	}
	return x
}

func IsRome(s string) (b bool) {

	b = false
	switch s {
	case "X":
		b = true
	case "I":
		b = true
	case "II":
		b = true
	case "III":
		b = true
	case "IV":
		b = true
	case "V":
		b = true
	case "VI":
		b = true
	case "VII":
		b = true
	case "VIII":
		b = true
	case "IX":
		b = true
	}
	return b
}

func ToDigit(s string) (x int) {
	switch s {
	case "10", "X":
		x = 10
	case "1", "I":
		x = 1
	case "2", "II":
		x = 2
	case "3", "III":
		x = 3
	case "4", "IV":
		x = 4
	case "5", "V":
		x = 5
	case "6", "VI":
		x = 6
	case "7", "VII":
		x = 7
	case "8", "VIII":
		x = 8
	case "9", "IX":
		x = 9
	}

	return x
}

func PrintRome(x int) {

	var s string = ""
	if x > 10 {
		switch x / 10 {
		case 9:
			s = "C"
		case 8:
			s = "LXXX"
		case 7:
			s = "LXX"
		case 6:
			s = "LX"
		case 5:
			s = "L"
		case 4:
			s = "XL"
		case 3:
			s = "XXX"
		case 2:
			s = "XX"
		case 1:
			s = "X"
		}
		x = x % 10
	}
	switch x {
	case 10:
		s += "X"
	case 1:
		s += "I"
	case 2:
		s += "II"
	case 3:
		s += "III"
	case 4:
		s += "IV"
	case 5:
		s += "V"
	case 6:
		s += "VI"
	case 7:
		s += "VII"
	case 8:
		s += "VIII"
	case 9:
		s += "IX"
	}
	fmt.Println(s)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b, c string
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	sub := strings.Split(text, " ")
	if len(sub) == 3 {

		var zn1, zn2 int

		a = sub[0]
		b = sub[1]
		c = sub[2]

		zn1 = ToDigit(a)

		zn2 = ToDigit(c)

		if zn1 > 10 || zn2 > 10 || a[0] == '-' || c[0] == '-' {
			fmt.Println("ERR: входное число болще десяти, алло")
			return
		}

		if IsRome(a) && IsRome(c) {
			if calcul(zn1, zn2, b) > 0 {

				PrintRome(calcul(zn1, zn2, b))

			} else {

				fmt.Println("ERR: в римской системе нет отрицательных чисел.")
			}
		} else if IsRome(a) || IsRome(c) {

			fmt.Println("ERR: одновременно используются разные системы счисления.")

		} else {

			fmt.Println(calcul(zn1, zn2, b))
		}
	} else if len(sub) < 3 {
		fmt.Println("ERR: строка не является математической операцией")
	} else {
		fmt.Println("ERR: формат математической операции не удовлетворяет заданию")
	}
}
