package main

func Match(r rune, i int, ascii map[byte][]string) string {
	var result string
	for ind, v := range ascii {
		if rune(ind) == r {
			result = v[i]
			break
		}
	}
	return result
}

func NewLine(tab []string) bool {
	for i := 0; i < len(tab); i++ {
		if tab[i] != "" {
			return false
		}
	}
	return true
}

func Printable(tab []rune) bool {
	for i := 0; i < len(tab); i++ {
		if tab[i] < 32 || tab[i] > 126 {
			return false
		}
	}
	return true
}

func Banner(s string) string {
	return "./" + s + ".txt"
}

func OutputName(S string) string {
	var filename string
	var j int
	for i := 0; i < len(S); i++ {
		if S[i] != '=' {
			j++
		} else {
			break
		}
	}
	for k := j + 1; k < len(S); k++ {
		filename += string(S[k])
	}
	return filename
}
