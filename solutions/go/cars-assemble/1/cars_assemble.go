package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    CalculateWorkingCarsPerHour := float64(productionRate) * float64(successRate/100)
    return CalculateWorkingCarsPerHour
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    CalculateWorkingCarsPerMinute := (float64(productionRate) * float64(successRate/100)) / 60
    return int(CalculateWorkingCarsPerMinute)
	panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	carsGroup := carsCount / 10
    remainingCars := carsCount % 10
    CalculateCost := (carsGroup * 95000) + (remainingCars * 10000)
    return uint(CalculateCost)
	panic("CalculateCost not implemented")
}
