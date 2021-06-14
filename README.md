# TCP_Chat_Server
Chat Server in Golang using channels, goroutines, and the `net` package that supports TCP. 
The TCP Chat application has the following components:

`client: current user and its connection `   
`server: which manages all incoming commands, and stores chat_rooms and clients  `  
`chat_room  `  
`command: from the client to the server  `  
`TCP server: to accept network connections  `  

# Available Commands
All commands begin with a slash '/'
 - `/username <username>` : Client can provide a username. If no username is provided, the user will be anonymous
 - `/join <chat_room>` : Client can join a chat_room. If the chat_room doesn't exist, a new chat_room will be created. 
 - `/rooms` : A list of all chat_rooms available to join will be displayed
 - `/send <message>` : Broadcasts message to everyone in the channel
 - `/quit` : exits from the chat server

 # To start the Chat server:
 - Open a terminal to run the server: `go run *.go`
 - Open other terminals to run clients: `telnet localhost 8888`

