/*
	Gumball API in Go (Version 2)
	Uses MongoDB and RabbitMQ 
*/
	
package main

func main() {
	server := NewServer()
	server.Run(":" + port)
}
