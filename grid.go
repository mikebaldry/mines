package main

import (
	"fmt"
)

type Grid struct {
	table   [][]*Tile
	rows    int
	columns int
}

func NewGrid(rows int, columns int) *Grid {
	result := make([][]*Tile, rows)
	for x := 0; x < rows; x++ {
		result[x] = make([]*Tile, columns)
		for y := 0; y < columns; y++ {
			result[x][y] = &Tile{
				Bomb: false,
				Hint: 0,
			}
		}
	}

	return &Grid{table: result, rows: rows, columns: columns}
}

func (g *Grid) ScatterBombs(bombs int) {
	for bombsRequired := bombs; bombsRequired > 0; {
		x := r.Intn(g.rows - 1)
		y := r.Intn(g.columns - 1)

		if !g.table[x][y].Bomb {
			g.table[x][y].Bomb = true
			bombsRequired--
		}
	}
}

func (g *Grid) UpdateHints() {
	for x := 0; x < g.rows; x++ {
		for y := 0; y < g.columns; y++ {
			g.table[x][y].Hint = 0
		}
	}

	incGridHint := func(x int, y int) {
		if x < 0 || x > g.rows-1 {
			return
		}

		if y < 0 || y > g.columns-1 {
			return
		}

		if g.table[x][y].Bomb {
			return
		}

		g.table[x][y].Hint++
	}

	for x := 0; x < g.rows; x++ {
		for y := 0; y < g.columns; y++ {
			tile := g.table[x][y]

			if tile.Bomb {
				incGridHint(x-1, y-1) // top left
				incGridHint(x, y-1)   // top
				incGridHint(x+1, y-1) // top right

				incGridHint(x-1, y) // left
				incGridHint(x+1, y) // right

				incGridHint(x-1, y+1) // bottom left
				incGridHint(x, y+1)   // bottom
				incGridHint(x+1, y+1) // bottom right
			}
		}
	}
}

func (g *Grid) revealHintsAndBlanks(x int, y int) {
	if x < 0 || x > g.rows-1 {
		return
	}

	if y < 0 || y > g.columns-1 {
		return
	}

	tile := g.table[x][y]
	if tile.Revealed {
		return
	}

	if tile.Bomb {
		return
	}

	tile.Revealed = true

	if tile.Hint > 0 {
		return
	}

	g.revealHintsAndBlanks(x-1, y-1) // top left
	g.revealHintsAndBlanks(x, y-1)   // top
	g.revealHintsAndBlanks(x+1, y-1) // top right

	g.revealHintsAndBlanks(x-1, y) // left
	g.revealHintsAndBlanks(x+1, y) // right

	g.revealHintsAndBlanks(x-1, y+1) // bottom left
	g.revealHintsAndBlanks(x, y+1)   // bottom
	g.revealHintsAndBlanks(x+1, y+1) // bottom right
}

func (g *Grid) Reveal(x int, y int) error {
	tile := g.table[x][y]

	if tile.Revealed {
		return fmt.Errorf("tile %d, %d is revealed already", x, y)
	}

	tile.Revealed = true

	if tile.Bomb {
		return fmt.Errorf("BOOM!")
	}

	if tile.Hint > 0 {
		return nil
	}

	g.revealHintsAndBlanks(x-1, y-1) // top left
	g.revealHintsAndBlanks(x, y-1)   // top
	g.revealHintsAndBlanks(x+1, y-1) // top right

	g.revealHintsAndBlanks(x-1, y) // left
	g.revealHintsAndBlanks(x+1, y) // right

	g.revealHintsAndBlanks(x-1, y+1) // bottom left
	g.revealHintsAndBlanks(x, y+1)   // bottom
	g.revealHintsAndBlanks(x+1, y+1) // bottom right

	return nil
}

func (g *Grid) Flag(x int, y int) error {
	tile := g.table[x][y]

	tile.Flagged = !tile.Flagged

	return nil
}

type Tile struct {
	Bomb     bool
	Hint     int
	Revealed bool
	Flagged  bool
}
