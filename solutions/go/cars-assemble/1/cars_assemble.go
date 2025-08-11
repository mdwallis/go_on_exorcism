package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    // productionRate: number of cars produced per hour
    var numCarsProducedPerMinute = float64(productionRate) / 60.0
    return int(numCarsProducedPerMinute * (successRate / 100.0))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	return uint(carsCount / 10) * uint(95000) + uint(carsCount % 10) * uint(10000)
}
