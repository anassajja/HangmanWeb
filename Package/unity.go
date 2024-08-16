package Package

import (
	utils "hangman/utils"
	"strings"
)

var word string

type GameState struct {
	SelectedWord   string
	ShadowWord     string
	GuessedLetters []string
	WrongLetters   []string
	TriesLeft      int
	WicheImage     int
}

// déclare Hangman de type GameState et l'initialise avec une nouvelle instance de la structure GameState.
var Hangman GameState = GameState{}

func GameStart(word string) {
	Hangman.SelectedWord = word
	Hangman.ShadowWord = strings.Repeat("_", len(word))
	Hangman.GuessedLetters = []string{}
	Hangman.WrongLetters = []string{}
	Hangman.TriesLeft = 10
	Hangman.WicheImage = 0
}

func Won() bool {
	return Hangman.ShadowWord == Hangman.SelectedWord // vérifier si le joueur a gagné la partie
}
func Lost() bool {
	return Hangman.TriesLeft <= 0 // vérifier si le joueur a perdu la partie
}

func GuessLetter(letter string) int { //prend letter de type string et qui retourne un entier.

	//vérifie si la variable letter est vide, contient un espace ou contient plus d'un caractère.
	// Si l'une de ces conditions est vraie, la fonction retourne -1, ce qui signifie que la lettre proposée est invalide.
	if letter == "" || letter == " " || len(letter) > 1 {
		return -1
	}
	//lettre proposée a déjà été devinée auparavant. Si c'est le cas, la fonction retourne 0,
	//ce qui signifie que la lettre a déjà été devinée.
	if utils.Contains(Hangman.GuessedLetters, letter) {
		return 0
	}
	//vérifie si la lettre proposée a déjà été proposée auparavant mais était incorrecte.
	//Si c'est le cas, la fonction retourne 0, ce qui signifie que la lettre a déjà été proposée et était incorrecte.
	if utils.Contains(Hangman.WrongLetters, letter) {
		return 0
	}
	//initialise une variable booléenne nommée found à false.
	//Cette variable sera utilisée plus tard pour indiquer si la lettre proposée a été trouvée dans le mot caché.
	found := false

	//initialise une variable new_shadow_word à une chaîne de caractères vide.
	//Cette variable sera utilisée pour stocker le nouveau mot caché après que le joueur ait proposé une lettre.
	new_shadow_word := ""

	//cela boucle sur chaque caractère du mot choisi (SelectedWord). Pour chaque caractère,
	//la boucle vérifie si la lettre proposée (letter) est égale au caractère.
	//Si c'est le cas, la lettre est ajoutée à la variable new_shadow_word.
	//Sinon, le caractère correspondant du mot caché actuel (ShadowWord) est ajouté à new_shadow_word.
	for idx, l := range Hangman.SelectedWord {
		if string(l) == letter {
			found = true
			new_shadow_word += letter
		} else {
			new_shadow_word += string(Hangman.ShadowWord[idx])
		}
	}
	Hangman.ShadowWord = new_shadow_word
	//vérifie si la lettre proposée a été trouvée dans le mot caché (found est true).
	//Si c'est le cas, la lettre est ajoutée à la liste des lettres devinées
	//(GuessedLetters) et la fonction retourne 1, ce qui signifie que la lettre était correcte.
	if found {
		Hangman.GuessedLetters = append(Hangman.GuessedLetters, letter)
		return 1
	}
	Hangman.TriesLeft -= 1
	Hangman.WicheImage += 1
	//ajoute la lettre proposée à la liste des lettres incorrectes (WrongLetters)
	Hangman.WrongLetters = append(Hangman.WrongLetters, letter)

	return -1
}
