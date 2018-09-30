package samp

//ConnectWithRCONPasswordAndServerPassword tries connceting to a server
//supplying IP, Port, Username, RCON password and server password.
func ConnectWithRCONPasswordAndServerPassword(username, address string, port int, rconPassword, serverPassword string) {
	connectWithRCONPasswordAndServerPassword(username, address, port, rconPassword, serverPassword)
}

//ConnectWithRCONPassword tries connceting to a server
//supplying IP, Port, Username and RCON password.
func ConnectWithRCONPassword(username, address string, port int, rconPassword string) {
	connectWithRCONPassword(username, address, port, rconPassword)
}

//ConnectWithServerPassword tries connceting to a server
//supplying IP, Port, Username and server password.
func ConnectWithServerPassword(username, address string, port int, serverPassword string) {
	ConnectWithServerPassword(username, address, port, serverPassword)
}
