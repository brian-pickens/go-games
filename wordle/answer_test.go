package wordle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Guess_Match_ResultShouldBeTrue(t *testing.T) {
	answer := newAnswer("abcde")
	result, _ := answer.guess(([5]rune)([]rune("abcde")))
	assert.True(t, result)
}

func Test_Guess_NoMatch_ResultShouldBeFalse(t *testing.T) {
	answer := newAnswer("bcdef")
	result, _ := answer.guess(([5]rune)([]rune("abcde")))
	assert.False(t, result)
}

func Test_Guess_Match_GuessStateShouldBeAllCorrect(t *testing.T) {
	word := "abcde"
	guess := "abcde"
	answer := newAnswer(word)
	expected := [5]state{correct, correct, correct, correct, correct}
	_, actual := answer.guess(([5]rune)([]rune(guess)))
	assert.ElementsMatch(t, expected, actual)
}

func Test_Guess_PartialMatch_GuessStateShouldMatch(t *testing.T) {
	word := "abccc"
	guess := "abcde"
	answer := newAnswer(word)
	expected := [5]state{correct, correct, correct, incorrect, incorrect}
	_, actual := answer.guess(([5]rune)([]rune(guess)))
	assert.ElementsMatch(t, expected, actual)
}

func Test_Guess_WrongOrder_GuessStateShouldShowPresent(t *testing.T) {
	word := "aeghi"
	guess := "abcde"
	answer := newAnswer(word)
	expected := [5]state{correct, incorrect, incorrect, incorrect, present}
	_, actual := answer.guess(([5]rune)([]rune(guess)))
	assert.ElementsMatch(t, expected, actual)
}

func Test_Guess_WithExtraMatchingLettersShouldShowIncorrect(t *testing.T) {
	word := "talcy"
	guess := "batty"
	answer := newAnswer(word)
	expected := [5]state{incorrect, correct, present, incorrect, correct}
	_, actual := answer.guess(([5]rune)([]rune(guess)))
	assert.ElementsMatch(t, expected, actual)
}

func Test_CountLetters_ReturnsMapOfCorrectLetterCounts(t *testing.T) {
	word := "batty"
	answer := newAnswer(word)
	expected := map[rune]int{
		'a': 1,
		'b': 1,
		't': 2,
		'y': 1,
	}
	actual := answer.letterCount()
	assert.EqualValues(t, expected, actual)
}
