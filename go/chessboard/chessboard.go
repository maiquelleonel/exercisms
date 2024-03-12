package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

func mapBoard(board Chessboard) (map[int]int, map[string]int, int, int) {
	ranks := make(map[int]int)
	files := make(map[string]int)
	totalSquare := 0
	totalOccupied := 0
	for file, _files := range board {
		for rank, occupied := range _files {
			if occupied {
				files[file]++
				ranks[rank]++
				totalOccupied++
			}
			totalSquare++
		}

	}
	return ranks, files, totalSquare, totalOccupied
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	_, totalFiles, _, _ := mapBoard(cb)
	return totalFiles[file]

}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	totalRanks, _, _, _ := mapBoard(cb)
	return totalRanks[rank-1]
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	_, _, totalPairs, _ := mapBoard(cb)
	return totalPairs
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	_, _, _, totalOccupied := mapBoard(cb)
	return totalOccupied
}
