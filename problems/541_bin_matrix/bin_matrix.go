// Package bin_matrix
// difficulty: medium
// link: https://leetcode.com/problems/01-matrix/
package bin_matrix

import "math"

func updateMatrix(mat [][]int) [][]int {
	changed := true
	for changed {
		changed = false
		for i := 0; i < len(mat); i++ {
			for j := 0; j < len(mat[i]); j++ {
				if mat[i][j] != 0 {
					dist := getMinNeighbour(mat, i, j) + 1
					if dist != mat[i][j] {
						mat[i][j] = dist
						changed = true
					}
				}
			}
		}
	}
	return mat
}

func getMinNeighbour(mat [][]int, x, y int) int {
	min := math.MaxInt
	if x < len(mat)-1 {
		if mat[x+1][y] < min {
			min = mat[x+1][y]
		}
	}
	if y < len(mat[0])-1 {
		if mat[x][y+1] < min {
			min = mat[x][y+1]
		}
	}
	if x > 0 {
		if mat[x-1][y] < min {
			min = mat[x-1][y]
		}
	}
	if y > 0 {
		if mat[x][y-1] < min {
			min = mat[x][y-1]
		}
	}
	return min
}
