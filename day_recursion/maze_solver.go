package day_recursion

type Point struct {
	x int
	y int
}

type SolveRules struct {
	maze  []string
	wall  string
	start Point
	end   Point
	seen  [][]bool
	path  []Point
}

func walk(rules *SolveRules) bool {
	dir := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	if rules.start.y < 0 || rules.start.y >= len(rules.wall) {
		return false
	}

	if rules.maze[rules.start.y][rules.start.x] == rules.wall[0] {
		return false
	}

	if rules.start.x == rules.end.x && rules.start.y == rules.end.y {
		rules.path = append(rules.path, rules.end)
		return true
	}

	if rules.seen[rules.start.y][rules.start.x] {
		return false
	}

	//3 recurse
	//pre
	rules.seen[rules.start.y][rules.start.x] = true
	rules.path = append(rules.path, rules.start)
	for i := range dir {
		x := dir[i][0]
		y := dir[i][1]
		rules.start.x = x
		rules.start.y = y
		if walk(rules) {
			return true
		}
	}
	//post
	rules.path = rules.path[:len(rules.path)-1]
	return false
}

func MazeSolver(rules SolveRules) []Point {
	var seen [][]bool
	var path []Point

	for i := range rules.maze {
		seen[i] = make([]bool, len(rules.maze[i]))
	}

	walk(&rules)

	return path

}
