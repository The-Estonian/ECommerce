package validate

import "fmt"

func ValidateUserSession(cookie string) string {
	// return database.GetUserSession(cookie)
	fmt.Println("Validated!")
	return "0"
}
