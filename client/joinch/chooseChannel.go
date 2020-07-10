package joinch

import (
	"bufio"
	"fmt"
	"os"

	"../utils"
)

func chooseChannel(channels []string) string {
	scanner := bufio.NewScanner(os.Stdin)
Top:
	fmt.Println("List of channels:")
	for _, c := range channels {
		fmt.Println(c)
	}
	fmt.Print("Which channel? ")
	scanner.Scan()
	channelName := scanner.Text()
	if utils.StringInSlice(channelName, channels) {
		return channelName
	}
	fmt.Println("That is not an available channel.")
	goto Top
}
