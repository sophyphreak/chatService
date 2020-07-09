package main

import (
<<<<<<< HEAD
	"./dm"
)

func main() {
	dm.RunDM("ballzie")
=======
	"./menu"
	"./name"
)

func main() {

	username := name.GetUsername()
	menu.Menu(username)
>>>>>>> c5d6b432cc86c4b12cc02ef9f8e25f1980a2c979
}
