func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		seen := make(map[byte]bool)
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.'{
				continue
			}
			if seen[board[i][j]] == true {
				return false
			}
			seen[board[i][j]] = true
		}
	}

	for i:=0; i< len(board); i++{
		seen := make(map[byte]bool)
		for j := 0; j<len(board[i]); j++ {
			if board[j][i] == '.'{
				continue
			}
			if seen[board[j][i]] == true {
				return false
			}
			seen[board[j][i]] = true
		}
	}

	for square :=0; square < 9; square++ {
		seen := make(map[byte]bool)
		for i:=0; i< 3; i++ {
			for j :=0; j<3; j++ {
				row := (square / 3) * 3 + i
                col := (square % 3) * 3 + j
				if board[row][col] == '.'{
					continue
				}
				if seen[board[row][col]] ==true {
					return false
				}
				seen[board[row][col]]= true
			}
		}
	}
	return true
}
