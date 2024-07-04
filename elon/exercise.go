package elon

import "fmt"

func (c *Car) Drive() {
	resultingBattery := c.battery - c.batteryDrain
	resultingDistance := c.distance + c.speed

	if resultingBattery > 0 {
		c.battery = resultingBattery
		c.distance = resultingDistance
	}
}

func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d", c.battery) + "%"
}

func (c Car) CanFinish(trackDistance int) bool {
	maxDistance := c.speed * (c.battery / c.batteryDrain)

	return maxDistance >= trackDistance
}

// TODO: define the 'Drive()' method

// TODO: define the 'DisplayDistance() string' method

// TODO: define the 'DisplayBattery() string' method

// TODO: define the 'CanFinish(trackDistance int) bool' method

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
