package wordle

type answer [COLUMNS]rune
type state int

const (
	empty state = iota
	incorrect
	present
	correct
)

func newAnswer(str string) *answer {
	an := ([5]rune)([]rune(str))
	return &answer{an[0], an[1], an[2], an[3], an[4]}
}

func (an *answer) guess(guess [COLUMNS]rune) (result bool, guessState [COLUMNS]state) {
	result = false
	answerCount := letterCount(([COLUMNS]rune)(*an))
	guessCount := map[rune]int{}

	if guess == ([COLUMNS]rune)(*an) {
		return true, [COLUMNS]state{correct, correct, correct, correct, correct}
	}

	for a := 0; a < COLUMNS; a++ {
		guessState[a] = incorrect
		if guess[a] == an[a] {
			guessState[a] = correct
			guessCount[guess[a]]++
		}
	}

	for a := 0; a < COLUMNS; a++ {
		guessCount[guess[a]]++
		if guessState[a] != correct && guessCount[guess[a]] <= answerCount[guess[a]] {
			guessState[a] = present
		}
	}

	return result, guessState
}

func letterCount(runes [COLUMNS]rune) (count map[rune]int) {
	count = map[rune]int{}
	for _, r := range runes {
		count[r]++
	}
	return count
}
