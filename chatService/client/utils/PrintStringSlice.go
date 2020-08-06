package utils

import (
	"fmt"
)

func PrintStringSlice(slice []string) {
	for _, val := range slice {
		fmt.Println(val)
	}
}
