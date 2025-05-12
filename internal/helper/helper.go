package helper

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/roniahmad/parking-app/app/vars"
)

var (
	seeded *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func genRandAlphabet(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seeded.Intn(len(charset))]
	}
	return string(b)
}

/*
Emulate parking censor to issue ticket timestamp
output sample
*/
func GenerateTimeStamp() string {
	return time.Now().Local().Format(time.RFC3339)
}

/*
Emulate parking censor to detect car color
output sample RED
*/
func GenerateRandomCardColor() string {
	return vars.Colors[seeded.Intn(len(vars.Colors))]
}

/*
Emulate parking censor to detect car plate number
output sample KA-01-HH-1234
*/
func GenerateRandomCarNumber() string {
	var charsAlphabet = vars.Alphabet
	var charsNumber = vars.Number

	return fmt.Sprintf("%s-%s-%s-%s",
		genRandAlphabet(2, charsAlphabet),
		genRandAlphabet(2, charsNumber),
		genRandAlphabet(2, charsAlphabet),
		genRandAlphabet(4, charsNumber),
	)
}
