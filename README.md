# TCP_Chat_Server
TCP Chat Server in Golang
# Available Commands
All commands begin with a slash '/'
 - `/username <username>` : Client can provide a username. If no username is provided, the user will be anonymous
 - `/join <channel_name>` : Client ca join a channel/group. If channel_name doesn't exist, a new channel_name will be created. 
 - `/channels` : A list of all channels available to join will be displayed
 - `/send <message>` : Broadcasts message to everyone in the channel
 - `/quit` : exits from the chat server

 # To start the Chat server:
 - Open a terminal to run the server: `go run *.go`
 - Open other terminals to run clients: `telnet localhost 8888`

