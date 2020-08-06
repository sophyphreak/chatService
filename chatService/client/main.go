package main

import (
	"./menu"
	"./name"
)

func main() {

	username := name.GetUsername()
	menu.Menu(username)
}
