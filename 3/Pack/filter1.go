package Pack

func NePlusMinus(s string) (b bool) {

	switch s {
	case "1":
		b = true
	case "2":
		b = true
	case "3":
		b = true
	case "4":
		b = true
	case "5":
		b = true
	case "6":
		b = true
	case "7":
		b = true
	case "8":
		b = true
	case "9":
		b = true
	case "10":
		b = true
	}
	return b
}

func ToDigit(s string) (x int) {
	switch s {
	case "10":
		x = 10
	case "1":
		x = 1
	case "2":
		x = 2
	case "3":
		x = 3
	case "4":
		x = 4
	case "5":
		x = 5
	case "6":
		x = 6
	case "7":
		x = 7
	case "8":
		x = 8
	case "9":
		x = 9
	}

	return x
}
