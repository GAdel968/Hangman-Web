package hangmanweb

import (
	"math/rand"
	"strings"
	"time"

	hc "./hangman-classic"
)

// Game holds the state of the current hangman game
type Game struct {
	Word             string   // Current word being guessed
	WordToFind       string   // The original word to find
	Attempts         int      // Number of attempts left
	TriedLetters     []string // Letters already tried
	HangmanPositions []string // ASCII representations of hangman positions
	GameStatus       string   // Current status of the game (playing, won, lost)
}

// InitGame initializes a new hangman game
func InitGame() Game {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Read words from file
	words, err := hc.ReadWordsFromFile("words.txt")
	if err != nil {
		// Fallback words if file can't be read
		words = []string{"golang", "hangman", "web", "development", "programming"}
	}

	// Select random word
	wordToFind := words[rand.Intn(len(words))]

	// Initialize the word with underscores
	word := strings.Repeat("_ ", len(wordToFind))

	// Initialize hangman positions
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

	// Create and return game
	return Game{
		Word:             word,
		WordToFind:       wordToFind,
		Attempts:         6,
		TriedLetters:     []string{},
		HangmanPositions: positions,
		GameStatus:       "playing",
	}
}

// ProcessGuess processes a letter guess and updates the game state
func ProcessGuess(game Game, letterInput string) Game {
	// Convert to lowercase
	letter := strings.ToLower(letterInput)

	// Check if letter has already been tried
	for _, l := range game.TriedLetters {
		if l == letter {
			return game
		}
	}

	// Add letter to tried letters
	game.TriedLetters = append(game.TriedLetters, letter)

	// Check if letter is in the word
	found := false
	wordRunes := []rune(game.WordToFind)
	displayWordRunes := []rune(strings.ReplaceAll(game.Word, " ", ""))

	for i, r := range wordRunes {
		if string(r) == letter {
			displayWordRunes[i] = r
			found = true
		}
	}

	// Update display word
	game.Word = ""
	for _, r := range displayWordRunes {
		game.Word += string(r) + " "
	}
	game.Word = strings.TrimSpace(game.Word)

	// Update attempts if letter was not found
	if !found {
		game.Attempts--
	}

	// Check game status
	if !strings.Contains(game.Word, "_") {
		game.GameStatus = "won"
	}

	if game.Attempts <= 0 {
		game.GameStatus = "lost"
	}

	return game
}
