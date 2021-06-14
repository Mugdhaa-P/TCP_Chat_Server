package main

import (
	"log"
	"net"
	"fmt"
	"strings"
	"errors"
)
type server struct{
	//Key: name of room, Value:pointer to room
	rooms map[string]*room
	//Channel to send all messages from all clients
	commands chan command
}

//helper function to initialize server: 
//function returns pointer to server
func newServer() *server{
	return &server{
		rooms: make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd:= range s.commands {
		switch cmd.id {
		case CMD_USERNAME:
			s.username(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.getRooms(cmd.client)
		case CMD_SEND:
			s.send(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("New client has connected: %s", conn.RemoteAddr().String())

	//Note: No need to initialize rooms in client struct as it is nil
	c := &client{
		conn: conn,
		username: "anonymous",
		//Send-only channel for commands
		commands: s.commands,
	}

	c.readInput()
}

func (s *server) username(c *client, args []string) {
	if len(args) < 2 {
		c.send("Username is required. usage: /username <username>")
		return
	}

	c.username = args[1]
	c.send(fmt.Sprintf("Hey there, %s", c.username))
}

func (s *server) join(c *client, args []string) {
	if len(args) < 2 {
		c.send("Room name is required. usage: /join <room_name>")
		return
	}

	roomName := args[1]

	r, ok := s.rooms[roomName]

	//If room does not exist, create one
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		//Add the room to server.rooms property
		s.rooms[roomName] = r
	}

	//Add current client to the list of members of the room
	r.members[c.conn.RemoteAddr()] = c

	//Since a client can only be in one room, we quit the current room
	s.quitCurrentRoom(c)
	//Assign new room to the client
	c.room = r

	//Broadcast message to notify other members of the room
	r.broadcast(c, fmt.Sprintf("We have a new member! Welcome %s!", c.username))

	///Send message to the client
	c.send(fmt.Sprintf("Welcome to %s!", roomName))
}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		oldRoom := s.rooms[c.room.name]
		//delete client from members list
		delete(s.rooms[c.room.name].members, c.conn.RemoteAddr())
		oldRoom.broadcast(c, fmt.Sprintf("%s has left the chat-room!", c.username))
	}
}

func (s *server) listRooms(c *client) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}

	c.send(fmt.Sprintf("Available chat-rooms are: %s", strings.Join(rooms, ", ")))
}

func (s *server) send(c *client, args []string) {
	//Check if client room is not nil
	if c.room == nil {
		c.err(errors.New("You must join the chat-room first!"))
		return
	}
	if len(args) < 2 {
		c.send("message is required, usage: /msg MSG")
		return
	}

	//Everything from index 1 and later in args is the content of the message. 
	//Construct the message with .Join() 
	msg := strings.Join(args[1:], " ")

	//Broadcast the message in the client's current chat room
	c.room.broadcast(c, c.username+": "+msg)
}

func (s *server) getRooms(c *client) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}

	c.send(fmt.Sprintf("Available chat-rooms: %s", strings.Join(rooms, ", ")))
}

func (s *server) quit(c *client) {
	log.Printf("Client has left the chat: %s", c.conn.RemoteAddr().String())

	s.quitCurrentRoom(c)

	c.send("Sad to see you go :(")
	c.conn.Close()
}

