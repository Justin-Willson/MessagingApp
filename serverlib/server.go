package serverlib

import (
    "fmt"    
)

func StartServer() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			os.Exit(1)
		}
}