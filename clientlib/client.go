package clientlib

import (
    "bufio"
	"fmt"
	"os"
	"strings"
	"net"
)

const (
    CONN_HOST = "localhost"
)

func StartClient() {
	//Start Listener
	l, err := net.Listen("tcp", ":0")	
    if err != nil {
        panic(err)
    }
	defer l.Close()
	go listen(l)

	register(l.Addr().(*net.TCPAddr).Port)	
	
	//Listen for input from user
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
	
		if strings.Compare("\\quit", text) == 0 {
		  break
		}

		sendMsg(text)
	}
}

func listen(l net.Listener) {
	for {
		_, err := l.Accept()
		if err != nil {
			fmt.Printf("Some error %v", err)
		}
	}
}

func register(port int) {
	return
}

func sendMsg(msg string) {
	return
}