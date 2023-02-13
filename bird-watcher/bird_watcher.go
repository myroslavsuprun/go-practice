package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, v := range birdsPerDay {
		total += v
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	daysInWeek := 7
	weekSliceIndex := (week - 1) * daysInWeek
	weeklyBirds := birdsPerDay[weekSliceIndex : weekSliceIndex+daysInWeek]

	total := 0
	for _, v := range weeklyBirds {
		total += v
	}

	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	// Counting the range for our for loop, which is the number of even numbers.
	var oddDayLen int = len(birdsPerDay) / 2

	// Iterating through the loop for each even number.
	for i := 0; i < oddDayLen; i++ {
		index := i + i
		birdsPerDay[index] += 1
	}

	return birdsPerDay
}
