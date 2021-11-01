package hangman

import (
	"strings"
)

type Game struct {
	State        string   // Game State
	Letters      []string // Letters in the word to find
	FoundLetters []string // good guesses
	UsedLetters  []string // Used Letters
	TurnsLeft    int      // remaining attemps
}

func New(turns int, word string) *Game {
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = " _ "
	}
	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}
	return g
}

func (g *Game) MakeAGuess(guess string) {

	guess = strings.ToUpper(guess)

	switch g.State {
	case "won", "lost":
		return
	}

	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"

	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"

		g.RevealLetter(guess)

		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	} else {
		g.State = "badGuess"
		g.LostTurn(guess)

		if g.TurnsLeft <=0 {
			g.State ="lost"
		}
	}
}
func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}

func (g *Game) LostTurn(guess string) {
	g.TurnsLeft--
	g.UsedLetters=append(g.UsedLetters, guess)
}

func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}
func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}
