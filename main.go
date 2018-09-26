package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type tile struct {
	bomb     bool
	hint     int
	revealed bool
}

var r *rand.Rand

func main() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println("Mines")

	grid := NewGrid(10, 10)
	grid.ScatterBombs(20)
	grid.UpdateHints()

	reader := bufio.NewReader(os.Stdin)

	for {
		printGrid(grid)
		fmt.Println("(Q)uit (R)eveal (F)lag X Y")
		str, _ := reader.ReadString('\n')
		str = strings.Trim(str, "\n ")

		parts := strings.SplitN(str, " ", 3)

		x, err := strconv.Atoi(parts[1])
		if err != nil {
			break
		}
		y, err := strconv.Atoi(parts[2])
		if err != nil {
			break
		}

		switch strings.ToLower(parts[0]) {
		case "q":
			fmt.Println("Bye!")
			return
		case "r":
			if err := grid.Reveal(x, y); err != nil {
				fmt.Println(err)
			}
		case "f":
			grid.Flag(x, y)
		}
	}
}

func printGrid(grid *Grid) {
	// fmt.Print("\033[H\033[2J\033[3J")
	fmt.Println("  0 1 2 3 4 5 6 7 8 9")
	for y := 0; y < grid.columns; y++ {
		fmt.Printf("%d ", y)
		for x := 0; x < grid.rows; x++ {
			tile := grid.table[x][y]

			if tile.Revealed {
				if tile.Bomb {
					fmt.Printf("B ")
				} else {
					if tile.Hint > 0 {
						fmt.Printf("%d ", tile.Hint)
					} else {
						fmt.Printf("  ")
					}
				}
			} else {
				if tile.Flagged {
					fmt.Printf("F ")
				} else {
					fmt.Printf("_ ")
				}
			}
		}

		fmt.Println()
	}
}
