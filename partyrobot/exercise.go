package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %s years old!", name, fmt.Sprint(age))
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	welcomeMessage := Welcome(name)
	tableMessage := fmt.Sprintf("\nYou have been assigned to table %03d. Your table is on the left, exactly %.1f meters from here.\nYou will be sitting next to %s.", table, distance, neighbor)

	return welcomeMessage + tableMessage
}
