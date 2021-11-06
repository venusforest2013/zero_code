package spiral_order

func SpiralOrder(arr [][]int) []int {
	//边界判断
	rows := len(arr)
	if rows == 0 {
		return nil
	}
	cols := len(arr[0])
	if cols == 0 {
		return nil
	}
	total := rows * cols
	//定义结果 rows,cols,total
	res := make([]int, 0)
	//定义中间变量 visited| row | col |direction| directionIndex

	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	directionIndex := 0
	row, col := 0, 0
	for i := 0; i < total; i++ {
		visited[row][col] = true
		res = append(res, arr[row][col])
		nextRow := row + direction[directionIndex][0]
		nextCol := col + direction[directionIndex][1]
		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols || visited[nextRow][nextCol] {
			directionIndex = (directionIndex + 1) % 4
		}
		row = row + direction[directionIndex][0]
		col = col + direction[directionIndex][1]

	}
	return res
}
