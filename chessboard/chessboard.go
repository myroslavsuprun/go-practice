package chessboard

// Files array which contains whether the square is taken or not.
type File []bool

// Chessboard map with values of files in a chessboard.
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var total int
	cbFile, exists := cb[file]

	// If the file does not exist, return 0.
	if !exists {
		return 0
	}

	for _, fileSquare := range cbFile {
		if fileSquare {
			total++
		}
	}

	return total
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	var total int
	for _, v := range cb {
		if v[rank-1] {
			total++
		}
	}

	return total
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var total int
	for _, v := range cb {
		for range v {
			total++
		}
	}

	return total
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var total int

	for _, file := range cb {
		for _, rank := range file {
			if rank {
				total++
			}
		}
	}
	return total
}
