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
			a = a[1 : len(a)-1]
			c = c[1 : len(c)-1]
			sub1 := strings.Split(a, "")

			for i := 0; i <= len(a)-1; i++ {

				sl := []string{sub1[i]}
				for _, value := range sl {

					if c == value {
						fmt.Println(value)
						var s string = " "
						s = sub1[i-1] + sub1[i+1]

						fmt.Println(s)
						return
					}

				}
			}

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

	} else if len(sub) == 1 || len(sub) == 2 || len(sub) > 5 {
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

			a = a[1 : len(a)-0]
			b = b[0 : len(b)-1]

			d = d[1 : len(d)-1]

			if b == d {

				var s string = " "
				s = a

				fmt.Println("\"" + s + " \"")
				return
			}
			if a == d {
				var s string = " "
				s = b

				fmt.Println("\" " + s + "\"")
				return
			}

		}

	}
	if len(sub) == 5 {

		a := sub[0]
		b := sub[1]
		c := sub[2]
		d := sub[3]
		e := sub[4]
		if c != "-" {
			fmt.Println(Err)
			return
		}
		if c == "-" && d == "\"" && e == "\"" {
			if len(a)+len(b) > 10 || Pack.NePlusMinus(a) || Pack.NePlusMinus(b) || d != "\"" || e != "\"" {
				fmt.Println(Err)
				return
			}

			var s string = " "
			s = a + b

			fmt.Println(s)
			return
		}

	}
}
