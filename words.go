package main

import (
	"os"
	"runtime"
	"strings"
)

// type of os
const (
	WINDOWS = iota
	MAC
	LINUX
	OTHER
)

func detectedOs() byte {
	osName := runtime.GOOS
	var TargetOs byte
	switch osName {
	case "windows":
		TargetOs = WINDOWS
	case "darwin":
		TargetOs = MAC
	case "linux":
		TargetOs = LINUX
	default:
		TargetOs = OTHER
	}
	return TargetOs
}

func writeData(data string) {
	file, err := os.Create("./words.txt")
	checkError(err)
	defer file.Close()
	file.WriteString(data)
}

//LoadWords: if your os will mac or linux loads words from /usr/share/dict else download words from
func LoadWords() bool {
	//TODO: check file exits or not
	targetOs := detectedOs()
	switch targetOs {
	case WINDOWS:
		//TODO download word from website
		return false
	case OTHER:
		//TODO download word from website
		return false
	case MAC:
		data, err := os.ReadFile("/usr/share/dict/words")
		checkError(err)
		writeData(string(data))
		return true
	case LINUX:
		data, err := os.ReadFile("/usr/share/dict/words")
		checkError(err)
		writeData(string(data))
		return true
	}

	return false
}

func readWords() []string {
	data, err := os.ReadFile("./words.txt")
	checkError(err)
	return strings.Split(string(data), "\n")
}

func ChoisWord(length int) string {
	var words []string
	for _, word := range readWords() {
		if len(word) == length {
			words = append(words, word)
		}
	}

	return GetRandomWord(words)

}
