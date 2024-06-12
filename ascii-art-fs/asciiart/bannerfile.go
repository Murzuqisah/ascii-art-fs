package asciiart

import (
	"fmt"
	"os"
)

// GetBannerFileFromArgs determines the filename from the command line arguments and returns it
func GetBannerFileFromArgs(args []string) string {
	if len(args) == 3 {
		if args[2] == "shadow" {
			return "shadow.txt"
		} else if args[2] == "thinkertoy" {
			return "thinkertoy.txt"
		} else if args[2] == "standard"{
			return "standard.txt"
		} else {
			fmt.Println(`Unidentified file, Usage : go run . "text" bannerfile`)
			os.Exit(0)
		}
	}
	if len(args) != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		os.Exit(0)
	}
	return "standard.txt"
}
