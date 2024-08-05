package leetmicrosoft

import "fmt"

type Coordinate struct {
	x int
	y int
}

type VisitMap = map[string]bool

func spiralOrder(matrix [][]int) []int {
	visits := make(VisitMap, 0)
	limits := Coordinate{
		x: len(matrix[0]) - 1,
		y: len(matrix) - 1,
	}

	totalElements := len(matrix[0]) * len(matrix)
	visitedValues := make([]int, 0)
	visitedValues = append(visitedValues, matrix[0][0])
	visitCount := 1

	direction := "right"
	currentPosition := Coordinate{x: 0, y: 0}
	markVisit(currentPosition, visits)
	for visitCount < totalElements {
		newPosition, err := navigate(direction, currentPosition, limits, visits)

		currentPosition = newPosition

		if err != nil {
			direction = changeDirection(direction)
		} else {
			visitCount++
			visitedValues = append(visitedValues, matrix[currentPosition.y][currentPosition.x])
		}
	}

	return visitedValues
}

func navigate(
	direction string,
	currentPosition Coordinate,
	limits Coordinate,
	visits VisitMap) (Coordinate, error) {

	newPosition := Coordinate{x: currentPosition.x, y: currentPosition.y}

	switch direction {
	case "right":
		newPosition.x = newPosition.x + 1
	case "left":
		newPosition.x = newPosition.x - 1
	case "up":
		newPosition.y = newPosition.y - 1
	case "down":
		newPosition.y = newPosition.y + 1
	}

	isVisited := checkVisited(newPosition, visits)

	if newPosition.y >= 0 &&
		newPosition.x >= 0 &&
		newPosition.y <= limits.y &&
		newPosition.x <= limits.x &&
		!isVisited {
		markVisit(newPosition, visits)

		return newPosition, nil
	}

	return currentPosition, fmt.Errorf("can't navigate anymore")
}

func (c *Coordinate) getKey() string {

	key := fmt.Sprintf("%d,%d", c.y, c.x)

	return key
}

func markVisit(coordinate Coordinate, visits VisitMap) {
	key := coordinate.getKey()

	visits[key] = true
}

func checkVisited(coordinate Coordinate, visits VisitMap) bool {
	key := coordinate.getKey()

	_, isPresent := visits[key]

	return isPresent
}

func changeDirection(direction string) string {
	switch direction {
	case "right":
		return "down"
	case "left":
		return "up"
	case "down":
		return "left"
	default:
		return "right"
	}
}
