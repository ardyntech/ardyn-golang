// A small package to help generate random strings and numbers

package helpers

import (
	"math/rand"
	"time"
)

//-------------------------------------------------------------

func GenerateRandomString(length int) string {

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var runes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)

	for i := range b {

		b[i] = runes[rand.Intn(len(runes))]

	}

	return string(b)

}

//-------------------------------------------------------------

func GenerateRandomNumber(min int, max int) int {

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	return rand.Intn((max - min + 1) + min)

}

//-------------------------------------------------------------
