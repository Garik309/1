package main

import (
	"3/Pack"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {

	Err := errors.New("ОШИБКА введено некорректное выражение")

	reader := bufio.NewReader(os.Stdin)
	var a, b, c string
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	sub := strings.Split(text, " ")

	var x int

	a = sub[0]
	b = sub[1]
	c = sub[2]

	if len(sub) == 3 {
		if b != "+" && b != "-" && b != "*" && b != "/" {
			fmt.Println(Err)
			return
		}

		if b == "+" {
			if len(a) > 10 || len(c) > 10 || Pack.NePlusMinus(a) || Pack.NePlusMinus(c) {
				fmt.Println(Err)
				return
			}

			a = a[:len(a)-1]
			c = c[1 : len(c)-0]

			var s string = " "
			s = a + c

			fmt.Println(s)
			return

		} else if b == "-" {
			if len(a) > 10 || len(c) > 10 || Pack.NePlusMinus(a) || Pack.NePlusMinus(c) {
				fmt.Println(Err)
				return
			}

			var s string = " "
			s = a

			fmt.Println(s)
			return

		} else if b == "/" {
			x = Pack.ToDigit(c)
			if len(a) > 10 || x > 10 || x < 1 || Pack.NePlusMinus(a) || x == 0 {
				fmt.Println(Err)
				return
			}
			z := len(a) - 2
			n := (z / x)
			n = z - n

			var s string = " "
			s = a
			s = s[1 : len(s)-n-1]

			fmt.Println("\"" + s + "\"")
			return

		} else if b == "*" {
			x = Pack.ToDigit(c)
			if len(a) > 10 || x > 10 || x < 1 || Pack.NePlusMinus(a) || x == 0 {
				fmt.Println(Err)
				return
			}

			a = a[1 : len(a)-1]

			var s string = " "
			s = strings.Repeat(a, x)

			fmt.Println(Pack.Print(s))
			return

		}

	} else if len(sub) == 1 || len(sub) == 2 || len(sub) >= 5 {
		fmt.Println(Err)
	}

	if len(sub) == 4 {

		a := sub[0]
		b := sub[1]
		c := sub[2]
		d := sub[3]
		if c != "-" {
			fmt.Println(Err)
			return
		}

		if c == "-" {
			if len(a)+len(b) > 10 || len(d) > 10 || Pack.NePlusMinus(a) || Pack.NePlusMinus(b) || Pack.NePlusMinus(d) {
				fmt.Println(Err)
				return
			}

			d = d[1 : len(d)-0]

			if b == d {

				var s string = " "
				s = a

				fmt.Println(s + " \"")
				return
			}
		}
	} else if len(sub) == 1 || len(sub) == 2 || len(sub) >= 5 {
		fmt.Println(Err)
	}

}
