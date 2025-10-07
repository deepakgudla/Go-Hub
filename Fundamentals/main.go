package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	green  = "\033[32m"
	yellow = "\033[33m"
	red    = "\033[31m"
	reset  = "\033[0m"
	bold   = "\033[1m"
)

func main() {
	examples := GetAllExamples()

	if len(os.Args) > 1 {
		arg := strings.ToLower(os.Args[1])

		switch arg {
		case "--list", "-l":
			fmt.Println(bold + yellow + "📘 Available Examples:" + reset)
			for _, ex := range examples {
				fmt.Printf("  • %s\n", ex.Name())
			}
			return

		default:
			for _, ex := range examples {
				if strings.ToLower(ex.Name()) == arg {
					fmt.Printf("%s🚀 Running: %s%s\n\n", green, ex.Name(), reset)
					ex.Run()
					fmt.Printf("%s✅ Completed: %s%s\n", green, ex.Name(), reset)
					return
				}
			}
			fmt.Printf("%s❌ Unknown example: '%s'%s\n", red, os.Args[1], reset)
			fmt.Println(yellow + "Use '--list' to see available examples." + reset)
			return
		}
	}

	fmt.Printf("%s🚀 Running All %d Go Fundamentals Examples%s\n\n", green, len(examples), reset)

	for i, ex := range examples {
		fmt.Printf("%s%d.%s %s\n", bold, i+1, reset, ex.Name())
		ex.Run()
		fmt.Println()
	}

	fmt.Println(green + "✅ All examples completed successfully!" + reset)
}
