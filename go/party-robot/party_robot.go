package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	msg := fmt.Sprintf("Welcome to my party, %s!", name)
	msg += fmt.Sprintf("You have been assigned to table %.3d. ", table)
	msg += fmt.Sprintf("Your table is %s, exactly %.1f meters from here.\n", direction, distance)
	msg += fmt.Sprintf("You will be sitting next to %s.", neighbor)
	return msg
}
