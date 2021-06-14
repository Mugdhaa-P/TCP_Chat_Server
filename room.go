package main

import "net"

type room struct{
	name string 
	//Key: network end point address, Value: Pointer to client
	members map[net.Addr]*client
}

func (r *room) broadcast(sender *client, msg string) {
	//Loop over all members of the room
	for addr, m := range r.members {
		//Broadcast message to everyone except the sender
		if sender.conn.RemoteAddr() != addr {
			m.send(msg)
		}
	}
}