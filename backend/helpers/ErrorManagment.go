package helpers

import (
	"fmt"
	"os"
)

func CheckErr(fromWhere string, err error) {
	if err != nil {
		// Log error to terminal
		fmt.Println(fromWhere, "->", err)

		// Log error to the file
		file, fileErr := os.OpenFile("./error_logging/error_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if fileErr != nil {
			fmt.Println("Failed to open error log file:", fileErr)
			return
		}
		defer file.Close()

		errorMessage := fmt.Sprintf("%s -> %s\n", fromWhere, err.Error())

		if _, writeErr := file.WriteString(errorMessage); writeErr != nil {
			fmt.Println("Failed to write to error log file:", writeErr)
		}
	}
}
