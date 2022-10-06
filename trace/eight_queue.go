package trace

import "fmt"

func EightQueueExample() {
	ret := solveNQueens(4)
	fmt.Println(ret)
	fmt.Println(len(ret))

}

func solveNQueens(n int) [][]string {
	var solutions [][]string
	solutions = [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := map[int]bool{}
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	backtrack(queens, n, 0, columns, diagonals1, diagonals2, &solutions)
	return solutions
}

func backtrack(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool, solutions *[][]string) {
	if row == n {
		board := generateBoard(queens, n)
		*solutions = append(*solutions, board)
		return
	}
	for i := 0; i < n; i++ {
		if columns[i] {
			continue
		}
		diagonal1 := row - i
		if diagonals1[diagonal1] {
			continue
		}
		diagonal2 := row + i
		if diagonals2[diagonal2] {
			continue
		}
		queens[row] = i
		columns[i] = true
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true
		backtrack(queens, n, row+1, columns, diagonals1, diagonals2, solutions)
		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
