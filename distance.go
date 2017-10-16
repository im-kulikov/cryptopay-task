package main

func partSearch(word string, words []string, wordLen int) int {
	distMin := wordLen

	for _, s := range words {
		if dist := levenshteinDistance(word, s); dist == 1 {
			return 1
		} else if distMin > dist {
			distMin = dist
		}
	}

	return distMin
}

func stepSearch(word string, step int, wordLen int) int {
	var (
		ok    bool
		words []string
	)

	if words, ok = vWords[step]; !ok {
		return wordLen
	}

	return partSearch(word, words, wordLen)
}

// Решение такое:
// - берём длину слова:
// - сначала проходим по схожей длине и вычисляем дистанцию по алгоритму Левенштейна
// - если результат нам не подходит - идём по шагам (слева / справа)
func getDistanceFast(word string, count int) {
	// берём длину слова:
	wordLen := len(word)
	min := wordLen

	// сначала проходим по схожей длине
	// и вычисляем дистанцию по алгоритму Левенштейна
	if dist := partSearch(word, vWords[wordLen], wordLen); dist == 1 {
		notify <- count
		return
	} else if min > dist {
		min = dist
	}

	// если результат нам не подходит - идём по шагам (слева / справа)
	step := 1
	for min > step {
		// Проверяем шаг вправо:
		if dist := stepSearch(word, wordLen+step, wordLen); dist == 1 {
			notify <- count
			return
		} else if min > dist {
			min = dist
		}

		// Проверяем шаг влево:
		if dist := stepSearch(word, wordLen-step, wordLen); dist == 1 {
			notify <- count
			return
		} else if min > dist {
			min = dist
		}

		step++
	}

	notify <- min * count
}
