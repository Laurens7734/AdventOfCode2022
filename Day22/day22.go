package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
)

func main() {
	input := utils.ReadFile("../Datafiles/day22.txt")
	board := make([][]int, 0)
	maxlen := 0
	for _, line := range input {
		if line == "" {
			break
		}
		if len(line) > maxlen {
			maxlen = len(line)
		}
	}
	for _, line := range input {
		if line == "" {
			break
		}
		row := make([]int, 0)
		for _, char := range line {
			num := 2
			if char == '#' {
				num = 1
			} else if char == '.' {
				num = 0
			}
			row = append(row, num)
		}
		for len(row) < maxlen {
			row = append(row, 2)
		}
		board = append(board, row)
	}
	Assignment1(board, input[len(input)-1])
	Assignment2(board, input[len(input)-1])
}

func Assignment1(board [][]int, instructions string) {
	x := 0
	y := 0
	direction := 0
	for i, c := range board[0] {
		if c == 0 {
			x = i
			break
		}
	}
	instructionList := readInstructions(instructions)
	for _, ins := range instructionList {
		if ins == "L" {
			direction += 3
			direction %= 4
		} else if ins == "R" {
			direction += 1
			direction %= 4
		} else {
			num, _ := strconv.ParseInt(ins, 0, 64)
			x, y = move(x, y, direction, int(num), board)
		}
	}
	fmt.Println(1000*(y+1) + (x+1)*4 + direction)
}

func Assignment2(board [][]int, instructions string) {
	boards := make(map[string][][]int)
	for i := 0; i*50 < len(board[0]); i++ {
		for j := 0; j*50 < len(board); j++ {
			cords := fmt.Sprintf("%v,%v", i, j)
			newboard := make([][]int, 0)
			for a := 50 * j; a < 50*(j+1); a++ {
				newboard = append(newboard, board[a][i*50:(i+1)*50])
			}
			boards[cords] = newboard
		}
	}
	x := 0
	y := 0
	direction := 0
	boardx := 1
	boardy := 0
	instructionList := readInstructions(instructions)
	for i, ins := range instructionList {
		if i%100 == 0 {
			fmt.Println("next")
		}
		if ins == "L" {
			direction += 3
			direction %= 4
		} else if ins == "R" {
			direction += 1
			direction %= 4
		} else {
			num, _ := strconv.ParseInt(ins, 0, 64)
			x, y, direction, boardx, boardy = move2(x, y, direction, int(num), boardx, boardy, boards)
		}
	}
	finalx := boardx*50 + x + 1
	finaly := boardy*50 + y + 1
	fmt.Println(1000*(finaly) + (finalx)*4 + direction)
}

func move2(x, y, dir, steps, boardx, boardy int, boards map[string][][]int) (int, int, int, int, int) {
	xmov, ymov := getMoveNums(dir)
	board := boards[fmt.Sprintf("%v,%v", boardx, boardy)]
	for i := 0; i < steps; i++ {
		nextx := x + xmov
		nexty := y + ymov
		if nextx >= 0 && nextx < 50 && nexty >= 0 && nexty < 50 {
			if board[nexty][nextx] == 1 {
				break
			}
			if board[nexty][nextx] == 0 {
				x = nextx
				y = nexty
			}
		} else {
			nbx, nby, ndir, nex, ney := newBoard(boardx, boardy, dir, x, y)
			newboard := boards[fmt.Sprintf("%v,%v", nbx, nby)]
			if newboard[ney][nex] == 0 {
				board = newboard
				x = nex
				y = ney
				boardx = nbx
				boardy = nby
				dir = ndir
				xmov, ymov = getMoveNums(dir)
			}
		}
	}
	return x, y, dir, boardx, boardy
}

func getMoveNums(dir int) (int, int) {
	xmov := 0
	ymov := 0
	switch dir {
	case 0:
		xmov++
	case 1:
		ymov++
	case 2:
		xmov--
	case 3:
		ymov--
	}
	return xmov, ymov
}

func newBoard(bx, by, dir, x, y int) (int, int, int, int, int) {
	if bx == 0 && by == 2 {
		if dir == 0 {
			return 1, 2, 0, 0, y
		} else if dir == 1 {
			return 0, 3, 1, x, 0
		} else if dir == 2 {
			return 1, 0, 0, 0, 49 - y
		} else {
			return 1, 1, 0, 0, x
		}
	}
	if bx == 0 && by == 3 {
		if dir == 0 {
			return 1, 2, 3, y, 49
		} else if dir == 1 {
			return 2, 0, 1, x, 0
		} else if dir == 2 {
			return 1, 0, 1, y, 0
		} else {
			return 0, 2, 3, x, 49
		}
	}
	if bx == 1 && by == 0 {
		if dir == 0 {
			return 2, 0, 0, 0, y
		} else if dir == 1 {
			return 1, 1, 1, x, 0
		} else if dir == 2 {
			return 0, 2, 0, 0, 49 - y
		} else {
			return 0, 3, 0, 0, x
		}
	}
	if bx == 1 && by == 1 {
		if dir == 0 {
			return 2, 0, 3, y, 49
		} else if dir == 1 {
			return 1, 2, 1, x, 0
		} else if dir == 2 {
			return 0, 2, 1, y, 0
		} else {
			return 1, 0, 3, x, 49
		}
	}
	if bx == 1 && by == 2 {
		if dir == 0 {
			return 2, 0, 2, 49, 49 - y
		} else if dir == 1 {
			return 0, 3, 2, 49, x
		} else if dir == 2 {
			return 0, 2, 2, 49, y
		} else {
			return 1, 1, 3, x, 49
		}
	}
	if bx == 2 && by == 0 {
		if dir == 0 {
			return 1, 2, 2, 49, 49 - y
		} else if dir == 1 {
			return 1, 1, 2, 49, x
		} else if dir == 2 {
			return 1, 0, 2, 49, y
		} else {
			return 0, 3, 3, x, 49
		}
	}
	return 0, 0, 0, 0, 0
}

func move(x, y, direction, steps int, board [][]int) (int, int) {
	xmov := 0
	ymov := 0
	switch direction {
	case 0:
		xmov++
	case 1:
		ymov++
	case 2:
		xmov--
	case 3:
		ymov--
	}
	for i := 0; i < steps; i++ {
		nextx, nexty := checkWrap(x+xmov, y+ymov, board)

		if board[nexty][nextx] == 2 {
			for board[nexty][nextx] == 2 {
				nextx, nexty = checkWrap(nextx+xmov, nexty+ymov, board)
			}
		}
		if board[nexty][nextx] == 1 {
			break
		}
		if board[nexty][nextx] == 0 {
			x = nextx
			y = nexty
		}
	}
	return x, y
}

func checkWrap(x, y int, board [][]int) (int, int) {
	if x < 0 {
		x = len(board[0]) - 1
	}
	if x > len(board[0])-1 {
		x = 0
	}
	if y < 0 {
		y = len(board) - 1
	}
	if y > len(board)-1 {
		y = 0
	}
	return x, y
}

func readInstructions(input string) []string {
	response := make([]string, 0)
	s := ""
	for _, char := range input {
		if char == 'L' || char == 'R' {
			response = append(response, s)
			response = append(response, string(char))
			s = ""
		} else {
			s += string(char)
		}
	}
	if s != "" {
		response = append(response, s)
	}
	return response
}
