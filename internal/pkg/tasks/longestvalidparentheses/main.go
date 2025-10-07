package longestvalidparentheses

func longestValidParentheses1(s string) int {
	if s == "" || len(s) == 1 {
		return 0
	}

	if s[0] == ')' {
		return longestValidParentheses(s[1:])
	}

	if s[len(s)-1] == '(' {
		return longestValidParentheses(s[0 : len(s)-1])
	}

	if s == "()" {
		return 2
	}

	if isValid(s) {
		return len(s)
	}

	lhsMaxLen := longestValidParentheses(s[1:])
	rhsMaxLen := longestValidParentheses(s[0 : len(s)-1])

	if lhsMaxLen > rhsMaxLen {
		return lhsMaxLen
	}

	return rhsMaxLen
}

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := range s {
		if s[i] == ')' {
			if len(stack) == 0 {
				return false
			}

			if stack[len(stack)-1] == '(' {
				stack = stack[0 : len(stack)-1]
			}
		} else {
			stack = append(stack, '(')
		}
	}

	return len(stack) == 0
}

func longestValidParentheses(s string) int {
	newSubstr := true
	maxLen := 0
	memoryStack := make([]bool, 0, len(s))

	findMax := func() {
		if len(memoryStack) != 0 {
			substrLen := 0
			for _, v := range memoryStack {
				if v {
					substrLen += 2
				} else {
					substrLen = 0
				}

				if substrLen > maxLen {
					maxLen = substrLen
				}
			}
		}
	}

	/**
	Если у нас новая подстрока
		Пытаемся найти новую максимальную длину
		Инициализируем наш стэк заново
		Объявляем что мы нашли новую подстроку
		Нам нужен стэк для маркировки найденных пар
	Иначе
		Если у нас открывающая скобка
			То пытаемся в стэке закрыть ей пару
			Если пара не найдена то считаем что у нас не валидная строка
		Иначе
			Добавляем в наш стэк новое false значение, которое в будущем нужно закрыть
	*/

	for i := range s {
		currentCh := s[i]
		if newSubstr {
			if currentCh == '(' {
				findMax()

				newSubstr = false
				memoryStack = memoryStack[0:0]
				memoryStack = append(memoryStack, false)
			}
		} else {
			if currentCh == ')' {
				braceFound := false
				for k := len(memoryStack) - 1; k > -1; k-- {
					if !memoryStack[k] {
						memoryStack[k] = true
						braceFound = true
						break
					}
				}

				if !braceFound {
					newSubstr = true
					continue
				}
			} else {
				memoryStack = append(memoryStack, false)
			}
		}
	}

	findMax()

	return maxLen
}
