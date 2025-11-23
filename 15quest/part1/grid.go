package main

import (
	"fmt"
)

// 0 = empty space
type Grid struct {
	height int
	width int
	walls int
	position point
	cave [][]int
}

func newGrid() *Grid {
	return &Grid{
		height: 1,
		width: 1,
		walls: 0,
		position: point{Y:0, X:0},
		cave: [][]int{
			{start},
		},
	}
}

func (g *Grid) printGrid() {
	fmt.Println()
	if g.height == 0 || g.width == 0 {return}
	fmt.Println("height:", g.height, "width:", g.width)
	fmt.Println("walls:", g.walls, "positon:", g.position)
	for _, row := range g.cave {
		line := ""
		for _, num := range row {
			switch num {
			case wall:
				line += "#"
			case start:
				line += "S"
			case end:
				line += "E"
			default:
				line += " "
			}
		}
		fmt.Printf("[%v]\n", line)
	}
}

func (g *Grid) up(move int) {
	if move <= 0 {return}
	if g.cave[g.position.Y][g.position.X] == start && g.walls > 0 {fmt.Printf("ERROR: cant make start position into a wall\n"); return}
	for step:=0; step<move; step++ {
		if g.position.Y == 0 { // create new row at top of grid
			newcave := make([][]int, 0, g.height+1)
			newrow := make([]int, g.width)
			newcave = append(newcave, newrow)
			for _, row := range g.cave {
				newrow = make([]int, g.width)
				copy(newrow, row)
				newcave = append(newcave, newrow)
			}
			g.cave = newcave
			g.height = len(g.cave)
			// check if a wall
			if g.cave[g.position.Y][g.position.X] == wall { // everything moved down, so new position is alread up
				fmt.Printf("WARNING: couldn't move up due to wall at: %v, %v\n", g.position.Y, g.position.X)
				return
			}
			if g.cave[g.position.Y][g.position.X] != start {
				g.cave[g.position.Y][g.position.X] = wall
				g.walls++
			} 
		} else {
			// check if a wall
			if g.cave[g.position.Y-1][g.position.X] == wall { // look up if no new row added
				fmt.Printf("WARNING: couldn't move up due to wall at: %v, %v\n", g.position.Y-1, g.position.X)
				return
			}
			// only change if positon is not start
			if g.cave[g.position.Y-1][g.position.X] != start {
				g.cave[g.position.Y-1][g.position.X] = wall
				g.walls++
			}
			g.position.Y--
		}
	}
}

func (g *Grid) down(move int) {
	if move <= 0 {return}
	if g.cave[g.position.Y][g.position.X] == start && g.walls > 0 {fmt.Printf("ERROR: cant make start position into a wall\n"); return}
	for step:=0; step<move; step++ {
		if g.position.Y+1 >= g.height { // create new row at bottom of grid
			newrow := make([]int, g.width)
			g.cave = append(g.cave, newrow)
			g.height = len(g.cave)
		}
		// check if a wall
		if g.cave[g.position.Y+1][g.position.X] == wall {
			fmt.Printf("WARNING: couldn't move down due to wall at: %v, %v\n", g.position.Y+1, g.position.X)
			return
		}
		// only change if positon is not start
		if g.cave[g.position.Y+1][g.position.X] != start {
			g.cave[g.position.Y+1][g.position.X] = wall
			g.walls++
		}
		g.position.Y++
	}
}

func (g *Grid) left(move int) {
	if move <= 0 {return}
	if g.cave[g.position.Y][g.position.X] == start && g.walls > 0 {fmt.Printf("ERROR: cant make start position into a wall\n"); return}
	for step:=0; step<move; step++ {
		if g.position.X <= 0 {
			//add new col to start of each inner arr
			g.width = len(g.cave[g.position.Y])+1
			for i, row := range g.cave {
				newrow := make([]int, 1, g.width)
				newrow = append(newrow, row...)
				g.cave[i] = newrow
			}
			// check if a wall
			if g.cave[g.position.Y][g.position.X] == wall {
				fmt.Printf("WARNING: couldn't move left due to wall at: %v, %v\n", g.position.Y, g.position.X)
				return
			}
			// only change if positon is not start
			if g.cave[g.position.Y][g.position.X] != start {
				g.cave[g.position.Y][g.position.X] = wall
				g.walls++
			}
			// position doesn't need changing
		} else {
			// check if a wall
			if g.cave[g.position.Y][g.position.X-1] == wall {
				fmt.Printf("WARNING: couldn't move left due to wall at: %v, %v\n", g.position.Y, g.position.X-1)
				return
			}
			// only change if positon is not start
			if g.cave[g.position.Y][g.position.X-1] != start {
				g.cave[g.position.Y][g.position.X-1] = wall
				g.walls++
			}
			g.position.X--
		}
	}
}

func (g *Grid) right(move int) {
	if move <= 0 {return}
	if g.cave[g.position.Y][g.position.X] == start && g.walls > 0 {fmt.Printf("ERROR: cant make start position into a wall\n"); return}
	for step:=0; step<move; step++ {
		if g.position.X+1 >= len(g.cave[g.position.Y]) {
			//add new col to end of each inner arr
			g.width = len(g.cave[g.position.Y])+1
			for i, _ := range g.cave {
				g.cave[i] = append(g.cave[i], 0) // add space to end
			}
			
		}
		// check if a wall
		if g.cave[g.position.Y][g.position.X+1] == wall {
			fmt.Printf("WARNING: couldn't move right due to wall at: %v, %v\n", g.position.Y, g.position.X+1)
			return
		}
		// only change if positon is not start
		if g.cave[g.position.Y][g.position.X+1] != start {
			g.cave[g.position.Y][g.position.X+1] = wall
			g.walls++
		}
		g.position.X++
	}
}

func (g *Grid) mapCave(commandSet []steps) {
	
	directions := []string{"up", "right", "down", "left"}
	looking := 0
	for _, command := range commandSet {
		// turn
		switch command.turn {
		case "L":
			looking--
		case "R":
			looking++
		}
		if looking < 0 {looking = len(directions)-1}
		if looking >= len(directions) {looking = 0}
		// move
		switch directions[looking] {
		case "up":
			g.up(command.step)
		case "right":
			g.right(command.step)
		case "down":
			g.down(command.step)
		case "left":
			g.left(command.step)
		}
	}
	// mark end
	g.cave[g.position.Y][g.position.X] = end

}

// returns the number of steps from start - end
func (g *Grid) getShortestPath() int {
	S, err := g.getPoint("S", start)
	if err != nil {fmt.Println(err)}
	E, err := g.getPoint("E", end)
	if err != nil {fmt.Println(err)}

	fmt.Println("S:", S, "E:", E) //

	return findShortestPath(S, E)
}

func (g *Grid) findShortestPath(start, end point) int {
	// recursivly search for the shortest ?
	// or calculate directions as the crow flies
		// try those first, 
		// if met will a wall, recurse the 2 directions
			// keep attempting to move closer to the end point

}

func (g *Grid) getPoint(find string, comp int) (point, error) {
	for i, _ := range g.cave {
		for j, num := range g.cave[i] {
			if num == comp {
				return point{
					Y: i,
					X: j,
				}, nil
			}
		}
	}

	return point{}, fmt.Errorf("WARNING: couldn't find %v", find)
}
