package main

import (
	"strconv"
	"strings"
)

func changeCase(w string) string {
	s := strings.Fields(w)
	for i := 0; i < len(s); i++ {
		if strings.HasPrefix(s[i], "(") && !strings.HasSuffix(s[i], ")") {
			a := s[i] + s[i+1]
			a = strings.Trim(a, "()")

			v := strings.Split(a, ",")
			num, _ := strconv.Atoi(v[1])

			m := i - num
			if m < 0 {
				m = 0
			}

			for j := m; j < i; j++ {
				switch v[0] {
				case "up":
					s[j] = strings.ToUpper(s[j])
				case "low":
					s[j] = strings.ToLower(s[j])
				case "cap":
					s[j] = strings.Title(strings.ToLower(s[j]))
				}
			}

			s = append(s[:i], s[i+2:]...)

		}
	}
	return strings.Join(s, " ")
}
