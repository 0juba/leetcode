package main

import "math"

/**
Задача

Генетическая строка может быть представлена строкой длиной 8 символов,
где каждый символ — это одна из букв 'A', 'C', 'G' или 'T'.

Предположим, что нужно исследовать мутацию от строки startGene к строке endGene,
где одна мутация определяется как изменение ровно одного символа в строке.

Например:
"AACCGGTT" → "AACCGGTA" — это одна мутация.

Существует также генетический банк bank, который хранит все допустимые варианты мутаций.
Ген может считаться допустимым, только если он присутствует в bank.

Даны две строки — startGene и endGene, — а также банк генов bank.
Необходимо вернуть минимальное количество мутаций, необходимых для превращения startGene в endGene.
Если такого преобразования не существует, вернуть -1.

Начальная строка считается допустимой, поэтому она может отсутствовать в bank.

⸻

Примеры

Пример 1:

Ввод:  startGene = "AACCGGTT", endGene = "AACCGGTA", bank = ["AACCGGTA"]
Вывод: 1

Пример 2:

Ввод:  startGene = "AACCGGTT", endGene = "AAACGGTA", bank = ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
Вывод: 2
*/

// minMutation https://leetcode.com/problems/minimum-genetic-mutation/
func minMutation(startGene string, endGene string, bank []string) int {
	stack := make([][2]int, 0, 1000)
	bank = append(bank, startGene)
	stack = append(stack, [2]int{len(bank) - 1, 0})

	visited := make([]bool, len(bank))
	for i := range bank {
		visited[i] = false
	}

	found := false
	result := math.MaxInt
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		prevGeneIdx, minLen := top[0], top[1]
		for i, gene := range bank {
			if bank[prevGeneIdx] == gene {
				continue
			}

			if diff(bank[prevGeneIdx], gene) == 1 {
				if gene == endGene {
					if result > minLen+1 {
						result = minLen + 1
						found = true
					}
				} else {
					if !visited[i] {
						stack = append(stack, [2]int{i, minLen + 1})
					}
					visited[i] = true
				}
			}
		}
	}

	if result == math.MaxInt && !found {
		return -1
	}

	return result
}

func diff(geneA, geneB string) int {
	count := 0
	for i := 0; i < len(geneA); i++ {
		if geneA[i] != geneB[i] {
			count++
		}
	}

	return count
}
