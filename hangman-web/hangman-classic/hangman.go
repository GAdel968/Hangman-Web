package hangman_classic

import (
	"bufio"
	"os"
	"strings"
)

func ReadWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func DisplayHangman(attempts int) string {
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
 /|\  |
      |
      |
=========`,
		`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`,
		`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
	}

	return positions[6-attempts]
}
