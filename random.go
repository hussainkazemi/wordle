package main

import (
	"math/rand"
	"time"
)

func getRandomNumber(maxSize int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(maxSize)
}

func GetRandomWord(words []string) string {
	return words[getRandomNumber(len(words))]
}
