package filereader

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
		content = AppendString(content, scanner.Text())
	}
	return content
}

func AppendString(slice []string, data string) []string {
	m := len(slice)
	n := m + 1
	if n > cap(slice) {
		newSlice := make([]string, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	slice[m] = data
	return slice
}
