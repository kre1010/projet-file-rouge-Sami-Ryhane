package main

import (
	"fmt"
	"strings"
)

type Joueur struct{
	Pseudo string `json:"pseudo"`
	Mot string `json:"mot"`
	Lettres []string `json:"lettres"`
	Essais int `json:"essais"`
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
func jeumot(s string) (string, error) {
	var j Joueur
	var nouveauJoueurPseudo string
	fmt.Println("Pseudo:")
	fmt.Scanln(&nouveauJoueurPseudo)
	j.Pseudo = nouveauJoueurPseudo
	j.Mot = s
	var trucla string 
	for i := 0 ; i <= len(s); i++ {
		trucla += "_"
	}
	fmt.Println(trucla,'\n')
	fmt.Println(etapesPendu[j.Essais])
	var lettre string
	fmt.Scan(&lettre)
	for i, value := range s {
		if lettre == string(value){
			runes := []rune(trucla)
			runes[i] = value
			lettre = string(runes)
			fmt.Println(lettre)
		} else if strings.Contains(s,lettre) != true{
			fmt.Println("La lettre n'est pas présente")
			j.Essais += 1
		}
	}
	return "", nil
}
