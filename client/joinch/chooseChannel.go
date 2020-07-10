package joinch

import (
	"bufio"
	"fmt"
	"os"

	"../utils"
)

func chooseChannel(channels []string) string {
	for {
		fmt.Println("List of channels:")
		for _, c := range channels {
			fmt.Println(c)
		}
		fmt.Print("Which channel? ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		channelName := scanner.Text()
		if utils.StringInSlice(channelName, channels) {
			return channelName
		}
		fmt.Println("That is not an available channel.")
	}
}
