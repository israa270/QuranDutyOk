package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// numbers
var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// RandNumbers randSeq Random String
func RandNumbers(n int) []int {
	// rand.Seed(time.Now().UnixNano())

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	b := make([]int, n)
	for i := range b {
		b[i] = numbers[r.Intn(len(numbers))]
	}
	return b
}

// ArrayIntToString convert array to string
func ArrayIntToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
