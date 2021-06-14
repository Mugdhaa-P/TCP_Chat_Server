package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
)
//
type client struct{
	conn net.Conn
	username string
	room *room 
	commands chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		//Parse msg to get the command 'cmd' and arguments 'args'
		//Example: "/join #general" => args = ["/join", "#general"]
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		//Send message to the channel
		//When messages are sent to the 'commands' channel, they will be received by server's run() function
		switch cmd {
		case "/username":
			c.commands <- command{
				id:     CMD_USERNAME,
				client: c,
				args:   args,
			}
		case "/join":
			c.commands <- command{
				id:     CMD_JOIN,
				client: c,
				args:   args,
			}
		case "/rooms":
			c.commands <- command{
				id:     CMD_ROOMS,
				client: c,
			}
		case "/send":
			c.commands <- command{
				id:     CMD_SEND,
				client: c,
				args:   args,
			}
		case "/quit":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
			}
		default:
			c.err(fmt.Errorf("Command not supported: %s", cmd))
		}
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("ERROR: " + err.Error() + "\n"))
}

func (c *client) send(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}