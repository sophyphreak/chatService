package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userName := strings.TrimSpace(scanner.Text())
	return userName
}
