package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"strings"
)

type Joueur struct {
	Pseudo   string   `json:"pseudo"`
	Mot      string   `json:"mot"`
	Lettres  []string `json:"lettres"`
	Essais   int      `json:"essais"`
	MotCache []string `json:"motcache"`
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

	// play game
	player.initStar()
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

	listWord := strings.Fields(string(fileData))
	randomIndexWord := rand.IntN(len(listWord))
	player.Mot = listWord[randomIndexWord]
	player.MotCache = strings.Split(strings.Repeat("_", len(player.Mot)), "")
	randomIndexShowLetter := rand.IntN(len(player.MotCache))
	player.MotCache[randomIndexShowLetter] = string(player.Mot[randomIndexShowLetter])
	return nil
}

func (player *Joueur) initStar() {
	for player.Essais <= 6 && strings.Contains(strings.Join(player.MotCache, ""), "_") {
		fmt.Println(etapesPendu[player.Essais], '\n')
		fmt.Println("Essais:", player.Essais)
		fmt.Println("Mot :", player.MotCache)
		fmt.Println("indiquer une lettre")
		var inputLetter string
		fmt.Scanln(&inputLetter)
		if len(inputLetter) > 1 {
			if inputLetter == player.Mot {
				fmt.Println("Vous avez trouvé le mot")

			}
		}
		if slices.Contains(player.Lettres, inputLetter) {
			fmt.Println("La lettre est déjà indiqué, veuillez indiquer une autre lettre")
			continue
		}
		player.Lettres = append(player.Lettres, inputLetter)

		if strings.Contains(player.Mot, inputLetter) {
			for index, value := range player.Mot {
				if string(value) == inputLetter {
					player.MotCache[index] = string(value)
				}
			}
		} else {
			player.Essais += 1
		}

	}

	if !strings.Contains(strings.Join(player.MotCache, ""), "_") {
		fmt.Println(player.MotCache)
		fmt.Println("Vous avez trouver le mot !")
	} else {
		fmt.Println("Vous avez perdu")
	}
}

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
