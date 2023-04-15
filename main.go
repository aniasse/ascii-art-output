package main

import (
	"fmt"
	"os"
	"strings"
)

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

func main() {
	ascii := make(map[byte][]string)
	var index byte = 32
	if len(os.Args) == 4 {
		banner := Banner(os.Args[3])
		file, err := os.ReadFile(banner)
		if err != nil {
			fmt.Println("Error : Not a ascci file in the repertory")
		}
		if os.Args[3] == "thinkertoy" {
			Split := strings.Split(string(file), "\r\n")
			for i := 1; i+8 < len(Split); i += 9 {
				ascii[index] = Split[i : i+8]
				index++
			}
		} else {
			Split := strings.Split(string(file), "\n")
			for i := 1; i+8 < len(Split); i += 9 {
				ascii[index] = Split[i : i+8]
				index++
			}
		}

	}

	tabascii := ascii
	if len(os.Args) == 4 {
		if len(os.Args[2]) != 0 {
			split := strings.Split(os.Args[2], "\\n")
			if NewLine(split) {
				split = split[:len(split)-1]
			}
			var res string
			for _, v := range split {
				tabrune := []rune(v)
				if Printable(tabrune) {
					for j := 0; j < 8; j++ {
						for i := 0; i < len(tabrune); i++ {
							res += Match(tabrune[i], j, tabascii)
						}
						if len(tabrune) != 0 {
							res = "\n"
						} else {
							res += "\n"
							break
						}
					}
					res += "\n"

					Filename := OutputName(os.Args[1])
					sortie, err := os.Create(Filename)
					if err != nil {
						fmt.Println(err)
						return
					}
					_, err = sortie.WriteString(res)
					if err != nil {
						fmt.Println(err)
						sortie.Close()
						return
					}
					err = sortie.Close()
					if err != nil {
						fmt.Println(err)
						return
					}
				} else {
					fmt.Println("Error : Non-displayable character !!!")
				}

			}

		}
	} else {
		fmt.Println("Error:\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	}

}
