package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	str, _ := os.ReadFile("File.txt")
	var str1 string
	var str2 []string
	var temp string
	var temp1 int
	var stock string
	var soluce []string
	// on passe le tableau de bytes reçu dans un string puis dans un tableau de string pour facilité la manipulation
	for i := 0; i < len(str); i++ {
		if str[i] == 13 {
			str2 = append(str2, str1)
			str1 = ""
			i++
		} else if str[i] != 13 {
			str1 += string(str[i])
		}
	}
	str2 = append(str2, str1)
	fmt.Println("Fragement 1 :", str2[0])           // on affiche le premier fragement
	fmt.Println("Fragement 2 :", str2[len(str2)-1]) // on affiche le 2ème fragment
	soluce = append(soluce, str2[0])
	soluce = append(soluce, str2[len(str2)-1])
	for i := 0; i < len(str2); i++ {
		if str2[i] == "before" {
			temp = str2[i+1] //	On stock la valeur de la valeur de l'emplacement dans une strings
			break
		}
	}
	//On a changer la valeur stocker en un int on peut mainteant l'utilisé pour trouver le prochain fragement
	for i, letter := range temp {
		if i > 0 {
			temp1 = temp1*10 + int(letter-48)
		} else {
			temp1 += int(letter - 48)
		}
	}

	fmt.Println("Fragement 3 :", str2[temp1-1])
	soluce = append(soluce, str2[temp1-1])

	for i := 0; i < len(str2); i++ {
		if str2[i] == "now" {
			stock = str2[i-1] //On stock la string
			break
		}
	}
	result := stock[1]                                       //on définit une variable qui prend pour valeur le 2ème bytes de la string stocké
	result1 := int(result) / int(stock[0]+stock[1]+stock[2]) //on divise la valeur de ce bytes par la valeur total des bytes du mots stocker le resultat nous donnez l'emplacement du dernier fragement
	fmt.Println("Fragement 4 :", str2[result1])
	soluce = append(soluce, str2[result1])
	fmt.Println("La solution de l'enigme est : ", soluce)

	nbr := rand.Int()
	fmt.Println(nbr)
}
