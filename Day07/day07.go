package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadFile("../Datafiles/day07.txt")
	createFolderStructure(input)
	calculateFilesizes("/")
	Assignment1()
	Assignment2()
}

type Folder struct {
	path       string
	subfolders []string
	filesize   int
	parrent    string
}

var folders = make(map[string]Folder)

func Assignment1() {
	totalsize := 0
	for _, f := range folders {
		if f.filesize <= 100000 {
			totalsize += f.filesize
		}
	}
	fmt.Println(totalsize)
}

func Assignment2() {
	totalspace := 70000000
	neededspace := 30000000
	usedspace := folders["/"].filesize
	tofree := neededspace - (totalspace - usedspace)
	smallest := totalspace
	for _, f := range folders {
		if f.filesize > tofree {
			if f.filesize < smallest {
				smallest = f.filesize
			}
		}
	}
	fmt.Println(smallest)
}

func calculateFilesizes(path string) int {
	folder := folders[path]
	for _, sub := range folder.subfolders {
		folder.filesize += calculateFilesizes(sub)
	}
	folders[path] = folder
	return folder.filesize
}

func createFolderStructure(input []string) {
	topfolder := Folder{
		path:       "/",
		subfolders: make([]string, 0),
		filesize:   0,
		parrent:    "",
	}
	folders["/"] = topfolder
	var currentfolder Folder
	for _, line := range input {
		if line[0] == '$' {
			currentfolder = processCommand(line[2:], currentfolder)
		} else {
			if line[:3] == "dir" {
				newfolder := Folder{
					path:       currentfolder.path + "/" + line[4:],
					subfolders: make([]string, 0),
					filesize:   0,
					parrent:    currentfolder.path,
				}
				currentfolder.subfolders = AppendString(currentfolder.subfolders, newfolder.path)
				folders[newfolder.path] = newfolder
			} else {
				fsize, _ := strconv.ParseInt(strings.Fields(line)[0], 0, 64)
				currentfolder.filesize += int(fsize)
			}
		}
	}
	folders[currentfolder.path] = currentfolder
}

func processCommand(command string, currentfolder Folder) Folder {
	if command[:2] == "ls" {
		return currentfolder
	}
	folders[currentfolder.path] = currentfolder
	direction := command[3:]
	if direction == "/" {
		return folders["/"]
	}
	if direction == ".." {
		if currentfolder.parrent == "" {
			return currentfolder
		} else {
			return folders[currentfolder.parrent]
		}
	} else {
		return folders[currentfolder.path+"/"+direction]
	}
}

func AppendString(slice []string, data string) []string {
	m := len(slice)
	n := m + 1
	if n > cap(slice) {
		newSlice := make([]string, n)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	slice[m] = data
	return slice
}
