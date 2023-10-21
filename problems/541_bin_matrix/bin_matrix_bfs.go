// Package bin_matrix
// difficulty: medium
// link: https://leetcode.com/problems/01-matrix/
package bin_matrix

// This solution:
// https://leetcode.com/problems/01-matrix/solutions/1369741/c-java-python-bfs-dp-solutions-with-picture-clean-concise-o-1-space/
func updateMatrixBFS(mat [][]int) [][]int {
	queue := [][]int{}
	DIR := []int{0, 1, 0, -1, 0}
	m, n := len(mat), len(mat[0])
	for i, row := range mat {
		for j, num := range row {
			if num == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1 // unprocessed
			}

		}
	}
	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			nx, ny := x+DIR[i], y+DIR[i+1]
			if nx < 0 || nx == m || ny < 0 || ny == n {
				continue
			}
			if mat[nx][ny] == -1 {
				mat[nx][ny] = mat[x][y] + 1
				queue = append(queue, []int{nx, ny})
			}
		}
	}
	return mat
}
