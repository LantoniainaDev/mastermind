package mind

import (
	"math/rand"
	"time"
)

func Generate(max, taille int) []int {
	// generer une liste de 4 chiffres differents de maniere aleatoire
	secret := make([]int, taille)

	generator := rand.New(rand.NewSource(time.Now().Unix()))

	secret2 := generator.Perm(max)

	for i := 0; i < taille; i++ {
		secret[i] = secret2[i] + 1
	}

	return secret
}

func Match(model []int, hint []int) [2]int {
	var match [2]int

	//unifier les couleurs
	unifyed := unify(hint)

	match[0] = countin(model, unifyed)
	match[1] = placing(model, hint)
	match[0] -= match[1]

	return match
}

func unify(target []int) []int {
	unifyed := make([]int, 0)

	for i := 0; i < len(target); i++ {
		var inside bool = false
		e := target[i]
		for j := 0; j < len(unifyed); j++ {
			if e == unifyed[j] {
				inside = true
				break
			}
		}

		if !inside {
			unifyed = append(unifyed, e)
		}
	}

	return unifyed
}

func countin(model []int, target []int) int {
	// compter les elements presents dans le model
	var count int

	for i := 0; i < len(target); i++ {
		e := target[i]
		for j := 0; j < len(model); j++ {
			if e == model[j] {
				count += 1
				break
			}
		}
	}

	return count
}

func placing(model []int, hint []int) int {
	var goodPlace int

	for j := 0; j < len(model); j++ {
		if model[j] == hint[j] {
			goodPlace += 1
		}
	}

	return goodPlace
}
