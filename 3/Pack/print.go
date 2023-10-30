package Pack

func Print(a string) string {
	var s string = ""
	if len(a) <= 40 {
		s = "\"" + a + "\""
	} else if len(a) > 40 {
		s = a[:len(a)-(len(a)-40)]
		s = "\"" + s + "...\""
	}
	return s
}
