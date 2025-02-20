package main

import (
	"fmt"
	hw "hangmanweb"
	"html/template"
	"net/http"
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

var currentGame Game

// Define template functions
var templateFuncs = template.FuncMap{
	"sub": func(a, b int) int {
		return a - b
	},
}

func main() {
	// Initialize game
	currentGame = hw.InitGame()

	// Setup static file server
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Setup routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hangman", hangmanHandler)
	http.HandleFunc("/restart", restartHandler)

	// Start server
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Create template with functions
	tmpl := template.Must(template.New("index.html").Funcs(templateFuncs).ParseFiles("templates/index.html"))

	// Execute template with game data
	tmpl.Execute(w, currentGame)
}

func hangmanHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Parse form data
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		// Get the letter from the form
		letter := r.FormValue("letter")

		// Process the guess
		if len(letter) > 0 {
			currentGame = hw.ProcessGuess(currentGame, letter)
		}
	}

	// Redirect back to home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func restartHandler(w http.ResponseWriter, r *http.Request) {
	// Reset the game
	currentGame = hw.InitGame()

	// Redirect back to home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
