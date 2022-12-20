package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(location string) []string {
	f, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var content []string
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	return content
}

func Contains[K comparable](slice []K, item K) bool {
	for _, ob := range slice {
		if ob == item {
			return true
		}
	}
	return false
}

func Abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}

func Max(a, b int) int {
	if b > a {
		return b
	} else {
		return a
	}
}
