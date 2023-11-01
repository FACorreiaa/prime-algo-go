package day_recursion

type Point struct {
	X, Y int
}

func walk(maze []string, wall string, start Point, end Point, seen [][]bool, path []Point) bool {
	dir := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	if start.Y < 0 || start.Y >= len(wall) {
		return false
	}

	if maze[start.Y][start.X] == wall[0] {
		return false
	}

	if start.X == end.X && start.Y == end.Y {
		path = append(path, end)
		return true
	}

	if seen[start.Y][start.X] {
		return false
	}

	//3 recurse
	//pre
	seen[start.Y][start.X] = true
	path = append(path, start)
	for i := range dir {
		X := dir[i][0]
		Y := dir[i][1]
		start.X = X
		start.Y = Y
		if walk(maze, wall, Point{start.X + X, start.Y + Y}, end, seen, path) {
			return true
		}
	}
	//post
	path = path[:len(path)-1]
	return false
}

// MazeSolver main function
func MazeSolver(maze []string, wall string, start Point, end Point) []Point {
	var seen [][]bool
	var path []Point

	for i := range maze {
		seen[i] = make([]bool, len(maze[i]))
	}

	walk(maze, wall, start, end, seen, path)

	return path

}

// DrawPath aux function
func DrawPath(data []string, path []Point) []string {
	var data2 = make([][]rune, len(data))
	for i, row := range data {
		data2[i] = []rune(row)
	}

	for _, p := range path {
		if p.Y >= 0 && p.Y < len(data2) && p.X >= 0 && p.X < len(data2[p.Y]) {
			data2[p.Y][p.X] = '*'
		}
	}

	result := make([]string, len(data))
	for i, row := range data2 {
		result[i] = string(row)
	}

	return result
}
