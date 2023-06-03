package wordle

type answer [COLUMNS]rune
type state int

const (
	empty state = iota
	incorrect
	present
	correct
)

func NewAnswer(str string) (answer) {
	an := ([5]rune)([]rune(str))
	return answer{an[0], an[1], an[2], an[3], an[4]}
}

func (an *answer) Guess(guess [COLUMNS]rune) (result bool, guessState [COLUMNS]state) {
	result = false
	if guess == ([COLUMNS]rune)(*an) {
		result = true
	}

	anCount := an.letterCount()
	loopCount := map[rune]int{}
	for a := 0; a < COLUMNS; a++ {
		if result {
			guessState[a] = correct
			continue
		}

		if guess[a] == rune(0) {
			guessState[a] = empty
			continue
		}

		if guess[a] == an[a] {
			guessState[a] = correct
			continue
		}

		for b := 0; b < COLUMNS; b++ {
			if guess[a] == an[b] && loopCount[guess[a]] < anCount[guess[a]] {
				guessState[a] = present
				loopCount[guess[a]]++
				break
			}
		}

		if (guessState[a] != present) {
			guessState[a] = incorrect
		}
	}

	return result, guessState
}

func (an *answer) letterCount() (count map[rune]int) {
	count = map[rune]int{}
	for _, r := range an {
		count[r]++
	}
	return count
}
