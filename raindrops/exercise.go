package raindrops

import "fmt"

type Sound struct {
	number int
	sound  string
}

func Convert(number int) string {
	soundList := [3]Sound{
		{number: 3, sound: "Pling"},
		{number: 5, sound: "Plang"},
		{number: 7, sound: "Plong"},
	}

	result := ""

	for _, item := range soundList {
		if number%item.number == 0 {
			result += item.sound
		}
	}

	if result == "" {
		return fmt.Sprintf("%d", number)
	}

	return result
}
