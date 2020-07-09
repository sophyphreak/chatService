package dm

// display list of usernames
// user inputs who they want to talk to
// display all dms between user and other user
// allow user to input more messages
//		- messages are sent to server
// sends a GET request to /messages/direct/{user}/{otherUser} every .25 seconds
//		- adds any messages that aren't already displayed
// allows user to quit to main menu by typing "quit" or something