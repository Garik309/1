package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// библиотеки мы просто у тебя стырили, чтобы считывание с терминала адекватно реализовать

// калькулятор, принимает три значения, возвращает результат
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

// проверка является ли число римким, возвращает true или false
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

// превращает и римские и арабские строчки тип int
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

// самая блин сложная функция
func PrintRome(x int) {

	var s string = "" // делаем пустую строку
	if x > 10 {
		switch x / 10 { // если число больше 10, то делим его на 10 и смотрим что осталось
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
		x = x % 10 // берём остаток от деления на 10, чтобы вписать единицы
	}
	switch x { // это выполняется и если число больше 10 и если меньше
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
	fmt.Println(s) // эта функция ничего не возвращает, она выводит результат на экран
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b, c string
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	sub := strings.Split(text, " ") // стырили у тебя строчки))
	if len(sub) == 3 {              // если длина массива == 3, то всё хорошо

		var zn1, zn2 int

		a = sub[0]
		b = sub[1]
		c = sub[2]

		zn1 = ToDigit(a) // делаем строчки цифрами

		zn2 = ToDigit(c)

		if zn1 > 10 || zn2 > 10 || a[0] == '-' || c[0] == '-' { // проверка на от 1 до 10, ещё и отрицательные числа смотрим типа "-2"
			fmt.Println("от одного до десяти ало")
			return // завершаем работу программы если числа не те
		}

		if IsRome(a) && IsRome(c) { // если оба римские, то делаем калькулятор и printrome

			if calcul(zn1, zn2, b) > 0 {

				PrintRome(calcul(zn1, zn2, b))

			} else { // если результат меньше 0, то ругаемся

				fmt.Println("В римских цифрах нет отрицательных сука")
			}
		} else if IsRome(a) || IsRome(c) { // иначе если только одна цифра римская, то ругаемся

			fmt.Println("АААШИПКА половина римская ааааааа вторая половина нет пиздец!!")

		} else { // если никто не римский, считаем и выводим

			fmt.Println(calcul(zn1, zn2, b))
		}
	} else if len(sub) < 3 { // если длина массива меньше трёх - ругаемся
		fmt.Println("Это не пример, это фигня какая-то")
	} else { // если в терминале что-то непонятное написано - ругаемся
		fmt.Println("ыыы два операнда один оператор бла бла бла")
	}
}
