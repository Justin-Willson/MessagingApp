package serverlib

import (
    "fmt"
	"net"
	"os"  
	"strings"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

type Client struct {
	hostname string
	port string
}

func StartServer() {
	clients := make([]Client, 0)

	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			os.Exit(1)
		}
	defer l.Close()
	fmt.Println("Running")

	for {
		conn, err := l.Accept()
		if err != nil {
			// handle error
		}
		go handleRequest(conn, &clients)
	}
}

func handleRequest(conn net.Conn, clientList *[]Client) {
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
	  fmt.Println("Error reading:", err.Error())
	}
	data := strings.Split(string(buf), ":")
	fmt.Println(data[0])


	//Register
	if data[0] == "ADD" {
		registerNewClient(conn, data[1], clientList)
	}

	//TODO: Transmit

	//TODO: Close
	
	conn.Close()
}

func registerNewClient(conn net.Conn, port string, clientList *[]Client) {
	//TODO: Make this thread safe
	addr := conn.RemoteAddr()
	data := strings.Split(addr.String(), ":")
	client := Client{data[0], port}
	*clientList = append(*clientList, client)
}