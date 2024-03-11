package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	acc := 0
	for _, v := range birdsPerDay {
		acc += v
	}
	return acc
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	return TotalBirdCount(birdsPerDay[7*(week-1) : 7*week])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}
	return birdsPerDay
}
