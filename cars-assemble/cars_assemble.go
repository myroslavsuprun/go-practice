package cars

// carsInBatch - how many cars should be in a batch for discount.
const carsInBatch int = 10

// batchOneCarCost - how much does a car cost if it is in a batch.
const batchOneCarCost int = 9500

// costPerCarBatch - how much a batch costs.
const costPerCarBatch int = carsInBatch * batchOneCarCost

// costPerSeparateCar - how much a car costs which is not a discount batch.
const costPerSeparateCar int = 10000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	successRate = percentageConversion(successRate)
	return float64(productionRate) * successRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	successRate = percentageConversion(successRate)
	successfulCarsPerHour := successRate * float64(productionRate)

	productionPerMinute := minuteConversion(successfulCarsPerHour)

	return int(productionPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// carsBatches - count total cars batches in the current calculation.
	carsBatches := carsCount / carsInBatch
	// batchesTotalCost - calculate expenses per current car batch.
	batchesTotalCost := carsBatches * costPerCarBatch

	// separated - cars count which are not in the current batch.
	separatedCars := carsCount % carsInBatch
	// separatedCarsCost - cost for cars which are not in the batch.
	separatedCarsCost := separatedCars * costPerSeparateCar

	// totalCost - return value with total cost for the carsCount
	var totalCost uint = uint(batchesTotalCost + separatedCarsCost)

	return totalCost
}

// percentageConversion - convert 100% into 1
func percentageConversion(percantage float64) float64 {
	return percantage / float64(100)
}

// minuteConversion - convert values per hour into per minute.
func minuteConversion(perHour float64) float64 {
	hourInMinutes := float64(60)

	return perHour / hourInMinutes
}
