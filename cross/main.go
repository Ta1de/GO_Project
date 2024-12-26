package main

import (
	"fmt"
	"log"
)

func main() {
	R, C, Cross := inputCross()
	fmt.Printf("Кроссворд размером: %d %d\n", R, C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			fmt.Printf("%c ", Cross[i][j])
		}
		fmt.Println()
	}
	result := decideCross(Cross, R, C)
	fmt.Println(result)
}

func inputCross() (R int, C int, Cross [][]rune) {
	fmt.Scanln(&R, &C)
	if R < 2 {
		log.Fatal("minimum length 2")
	}

	Cross = make([][]rune, R)
	for i := 0; i < R; i++ {
		Cross[i] = make([]rune, C)
	}

	for i := 0; i < R; i++ {
		var row string
		fmt.Scanln(&row)
		for j := 0; j < C; j++ {
			Cross[i][j] = rune(row[j])
		}
	}

	return R, C, Cross
}

func decideCross(Cross [][]rune, R int, C int) string {
	var result string

	for i := 0; i < len(Cross); i++ {
		for j := 0; j < len(Cross[0]); j++ {
			if Cross[i][j] != '#' {
				if path := findStraightPath(Cross, i, j, R); path != "" {
					if result == "" || path < result {
						result = path
					}
				}
			}
		}
	}
	return result
}

func findStraightPath(Cross [][]rune, row, col, length int) string {
	directions := [4][2]int{
		{-1, 0}, // вверх
		{1, 0},  // вниз
		{0, -1}, // влево
		{0, 1},  // вправо
	}

	for _, dir := range directions {
		if path := canMoveStraight(Cross, row, col, length, dir[0], dir[1]); path != "" {
			return path
		}
	}
	return ""
}

func canMoveStraight(Cross [][]rune, row, col, length, rowDir, colDir int) string {
	var path []rune
	for step := 0; step < length; step++ {
		newRow := row + step*rowDir
		newCol := col + step*colDir

		if newRow < 0 || newRow >= len(Cross) || newCol < 0 || newCol >= len(Cross[0]) || Cross[newRow][newCol] == '#' {
			return ""
		}
		path = append(path, Cross[newRow][newCol])
	}
	return string(path)
}
