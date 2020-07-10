package addch

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"../utils"
)

// Begin is the entrypoint for the add channel functionality
func Begin() {
	fmt.Print("Enter new channel name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	channelName := scanner.Text()

	values := map[string]string{"name": channelName}
	jsonValue, _ := json.Marshal(values)
	_, err := http.Post(utils.URL+"/channel", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		fmt.Println("Was not able to reach server")
		return
	}
	fmt.Println("Your channel", channelName, "was created successfully!")
	return
}
