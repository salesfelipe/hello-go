package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	productionRateFl := float64(productionRate)

	return productionRateFl * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	productionPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)

	return int(productionPerHour / 60)
}

func calculateCostMod10(carsCount int) uint {
	if carsCount >= 10 {
		var quotient int = carsCount / 10

		return uint(quotient * 95000)
	}

	return 0
}

func calculateCostNotMod10(carsCount int) uint {
	if carsCount%10 != 0 {
		return uint(carsCount%10) * 10000
	}

	return 0
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

	return calculateCostNotMod10(carsCount) + calculateCostMod10(carsCount)
}
