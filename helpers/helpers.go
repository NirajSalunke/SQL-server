package helpers

import (
	"fmt"
	"regexp"
	"strings"
)

var reset string = "\033[0m"

func PrintGreen(text string) {
	green := "\033[32m"
	fmt.Println(green + text + reset)
}

func PrintRed(text string) {
	red := "\033[31m"
	fmt.Println(red + text + reset)
}

func CleanSQLQuery(rawSQL string) string {
	re := regexp.MustCompile("(?i)```sql\\s*|```") // Removes ```sql and ```
	return strings.TrimSpace(re.ReplaceAllString(rawSQL, ""))
}
