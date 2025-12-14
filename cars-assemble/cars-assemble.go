package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(workingPerHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const individualCost = 10000
	const groupCost = 95000 // 10 cars cost 95,000

	groupsOfTen := carsCount / 10
	remainingCars := carsCount % 10

	total := groupsOfTen*groupCost + remainingCars*individualCost
	return uint(total)
}
