package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintResult(duplicates map[string][]string) {
	color.Yellow("\n⚠️  Found %d sets of duplicates:\n", len(duplicates))

	for _, paths := range duplicates {
		color.White("Duplicate Set: ")
		for _, path := range paths {
			fmt.Printf(" 📄 %s\n", path)
		}

		fmt.Println("-----------------------------------------")
	}

}