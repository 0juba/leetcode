package wildcardmatching

import "strings"

func isMatch(s string, p string) bool {
	p = compressAsterisks(p)

	var i, j int
	for i < len(s) && j < len(p) {
		if p[j] == '?' || s[i] == p[j] {
			i++
			j++
			continue
		} else if p[j] == '*' {
			stopChar := byte(0)
			if j+1 < len(p) {
				stopChar = p[j+1]
			} else {
				// Это конец шаблона, дальше может быть что угодно
				return true
			}

			// просто нужно промотать до оставшейся длины шаблона
			if stopChar == '?' {
				for ; i < len(s) && len(s)-i > len(p)-j-1; i++ {
				}
				continue
			}

			// мне нужно найти стоп символ для текущей звездочки
			// стоп символ нужно искать с конца
			// начальный индекс будет равен оставшейся длине шаблона
			// если начальный индекс больше чем сама строка - то проверка пройдена
			found := false
			for k := len(p) - j - 1; k > -1 && len(s) > k; k-- {
				if s[k] == stopChar {
					i = k
					found = true
					break
				}
			}

			if !found {
				return false
			}

			j++
			continue
		}

		return false
	}

	return len(s) == i && len(p) == j
}

func compressAsterisks(s string) string {
	buf := strings.Builder{}
	prev := byte(0)
	for i := range s {
		if s[i] == '*' {
			if prev != '*' {
				buf.WriteByte(s[i])
			}
		} else {
			buf.WriteByte(s[i])
		}

		prev = s[i]
	}

	return buf.String()
}

func isMatchQuestion(s string, p string) bool {
	if len(s) != len(p) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] != p[i] && p[i] != '?' {
			return false
		}
	}

	return true
}
