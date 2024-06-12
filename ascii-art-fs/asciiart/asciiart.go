package asciiart

import (
	"fmt"
	"os"
	"strings"
)

// printing the string based on the ascii character range
func PrintASCIIArt(lines []string, arguments string) {
	argument := strings.Split(arguments, "\n")
	count := 0
	for _, arg := range argument {
		for _, chr := range arg {
			if chr < 32 || chr > 126 {
				fmt.Println("Error : Non ascii/printable characters found")
				os.Exit(0)
			}
		}
		if arg == "" {
			count++
			if count < len(argument) {
				fmt.Println()
			}
		} else {
			for i := 0; i < 8; i++ {
				for _, value := range arg {
					start := int(value-32)*9 + 1
					fmt.Print(lines[start+i])
				}
				fmt.Println()
			}
			
		}
	}
}
