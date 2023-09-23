package randomgen

import (
	"math/rand"
	"time"
)

const defaultLowerCaseCharset = "abcdefghijklmnopqrstuvwxyz"
const defaultUpperCaseCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const defaultNumericCharset = "0123456789"
const defaultAlphaNumbericCharset = defaultLowerCaseCharset + defaultUpperCaseCharset + defaultNumericCharset

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func LowerCaseString(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = defaultLowerCaseCharset[seededRand.Intn(len(defaultLowerCaseCharset))]
	}

	return string(b)
}

func UpperCaseString(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = defaultUpperCaseCharset[seededRand.Intn(len(defaultUpperCaseCharset))]
	}

	return string(b)
}

func NumericString(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = defaultNumericCharset[seededRand.Intn(len(defaultNumericCharset))]
	}

	return string(b)
}

func AlphaNumericString(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = defaultAlphaNumbericCharset[seededRand.Intn(len(defaultAlphaNumbericCharset))]
	}

	return string(b)
}
