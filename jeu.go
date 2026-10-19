package main

import (
	"fmt"
	"os"
	"strings"
)

type Joueur struct {
	Pseudo  string   `json:"pseudo"`
	Mot     string   `json:"mot"`
	Lettres []string `json:"lettres"`
	Essais  int      `json:"essais"`
}

var etapesPendu = []string{
	`
  +---+
      |
      |
      |
     ===`,
	`
  +---+
  O   |
      |
      |
     ===`,
	`
  +---+
  O   |
  |   |
      |
     ===`,
	`
  +---+
  O   |
 /|   |
      |
     ===`,
	`
  +---+
  O   |
 /|\  |
      |
     ===`,
	`
  +---+
  O   |
 /|\  |
 /    |
     ===`,
	`
  +---+
  O   |
 /|\  |
 / \  |
     ===`,
}

// Lancement jeu et gestion des étapes (init pseudo, selct word, game)
func InitGame() {
	var player Joueur

	// Pseudo
	player.initPeseudo()

	// Select word
	player.initWord()

	fmt.Println(player)
}

func (player *Joueur) initPeseudo() {
	fmt.Println("Pseudo:")
	fmt.Scanln(&player.Pseudo)
	// Validation pseudo....
}

func (player *Joueur) initWord() error {
	fileData, fileErr := os.ReadFile("./mot.txt")
	if fileErr != nil {
		fmt.Println("erreur")
		return fileErr
	}

	fmt.Println(string(fileData))
	listWord := strings.Fields(string(fileData))

	fmt.Println(listWord)

	/* j.Mot = s
	var trucla strings.Builder
	for i := 0; i <= len(s); i++ {
		trucla.WriteString("_")
	}
	fmt.Println(trucla.String(), '\n')
	fmt.Println(etapesPendu[j.Essais])
	var lettre string
	fmt.Scan(&lettre)
	for i, value := range s {
		if lettre == string(value) {
			runes := []rune(trucla.String())
			runes[i] = value
			lettre = string(runes)
			fmt.Println(lettre)
		} else if strings.Contains(s, lettre) != true {
			fmt.Println("La lettre n'est pas présente")
			j.Essais += 1
		}
	}
	return "", nil */
	return nil
}
