package hangmanweb

import (
	"math/rand"
	"strings"
	"time"

	hc "./hangman-classic"
)

type Game struct {
	Word             string
	WordToFind       string
	Attempts         int
	TriedLetters     []string
	HangmanPositions []string
	GameStatus       string
}

func InitGame() Game {
	rand.Seed(time.Now().UnixNano())

	words, err := hc.ReadWordsFromFile("words.txt")
	if err != nil {
		words = []string
	}

	wordToFind := words[rand.Intn(len(words))]

	word := strings.Repeat("_ ", len(wordToFind))

	positions := []string{
		`
  +---+
  |   |
      |
      |
      |
      |
=========`,
		`
  +---+
  |   |
  O   |
      |
      |
      |
=========`,
		`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`,
		`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`,
		`
  +---+
  |   |
  O   |
 /|\\  |
      |
      |
=========`,
		`
  +---+
  |   |
  O   |
 /|\\  |
 /    |
      |
=========`,
		`
  +---+
  |   |
  O   |
 /|\\  |
 / \\  |
      |
=========`,
	}

	return Game{
		Word:             word,
		WordToFind:       wordToFind,
		Attempts:         6,
		TriedLetters:     []string{},
		HangmanPositions: positions,
		GameStatus:       "playing",
	}
}

func ProcessGuess(game Game, letterInput string) Game {
	letter := strings.ToLower(letterInput)

	for _, l := range game.TriedLetters {
		if l == letter {
			return game
		}
	}

	game.TriedLetters = append(game.TriedLetters, letter)

	found := false
	wordRunes := []rune(game.WordToFind)
	displayWordRunes := []rune(strings.ReplaceAll(game.Word, " ", ""))

	for i, r := range wordRunes {
		if string(r) == letter {
			displayWordRunes[i] = r
			found = true
		}
	}

	game.Word = ""
	for _, r := range displayWordRunes {
		game.Word += string(r) + " "
	}
	game.Word = strings.TrimSpace(game.Word)

	if !found {
		game.Attempts--
	}

	if !strings.Contains(game.Word, "_") {
		game.GameStatus = "won"
	}

	if game.Attempts <= 0 {
		game.GameStatus = "lost"
	}

	return game
}
