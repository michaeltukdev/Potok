package prompt

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Input(promptText string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(promptText)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read input: %v", err)
		}
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Input cannot be empty.")
			continue
		}
		return input
	}
}
