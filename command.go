package main

type commandID int

const (
	//Enum type with strings
	CMD_USERNAME commandID = iota 
	CMD_JOIN
	CMD_ROOMS
	CMD_SEND
	CMD_QUIT
)

type command struct{
	id commandID
	client *client //Sender of current command
	args []string
}