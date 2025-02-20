package main

import (
	"fmt"
	"html/template"
	"net/http"

	hw "github.com/GAdel968/Hangman-Web/hangmanweb"
)

type Game struct {
	Word             string
	WordToFind       string
	Attempts         int
	TriedLetters     []string
	HangmanPositions []string
	GameStatus       string
}

var currentGame Game

var templateFuncs = template.FuncMap{
	"sub": func(a, b int) int {
		return a - b
	},
}

func main() {
	currentGame = hw.InitGame()

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hangman", hangmanHandler)
	http.HandleFunc("/restart", restartHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index.html").Funcs(templateFuncs).ParseFiles("templates/index.html"))

	tmpl.Execute(w, currentGame)
}

func hangmanHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		letter := r.FormValue("letter")

		if len(letter) > 0 {
			currentGame = hw.ProcessGuess(currentGame, letter)
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func restartHandler(w http.ResponseWriter, r *http.Request) {
	currentGame = hw.InitGame()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
