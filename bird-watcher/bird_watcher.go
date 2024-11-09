package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var count int = 0
	for index := 0; index < len(birdsPerDay); index++ {
		count += birdsPerDay[index]
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	targetWeekDays := birdsPerDay[week*7-7 : week*7]
	return TotalBirdCount(targetWeekDays)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index := 0; index < len(birdsPerDay); index += 2 {
		birdsPerDay[index]++
	}

	return birdsPerDay
}
