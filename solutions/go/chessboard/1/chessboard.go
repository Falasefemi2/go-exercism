package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	// panic("Please implement CountInFile()")
	count := 0
	f := cb[file]
	for _, sqaure := range f {
		if sqaure {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    rank = rank - 1
	for _, file := range cb {
		if rank < 0 || rank >= len(file) {
			return 0
		}
		break
	}

	count := 0
	for _, file := range cb {
		if file[rank] {
			count++
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, c := range cb {
		for range c {
			count++
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, c := range cb {
		for _, f := range c {
			if f {
				count++
			}
		}
	}
	return count
}
