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
		content = AppendSlice(content, scanner.Text())
	}
	return content
}

func AppendSlice[K any](slice []K, data K) []K {
	m := len(slice)
	n := m + 1
	if n > cap(slice) {
		newSlice := make([]K, n)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	slice[m] = data
	return slice
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
