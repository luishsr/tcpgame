package main

import (
	"fmt"

	tcp "github.com/luishsr/tcpconnector"
	host "github.com/luishsr/tcphost"
)

func main() {
	fmt.Println("Client")

	//tcp.Register("locahost:8001", "Luis", "Client")
	host.RunListener("localhost", "8002")

}
