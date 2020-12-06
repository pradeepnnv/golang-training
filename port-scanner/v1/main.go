// Reference: https://medium.com/devops-dudes/how-to-write-a-tcp-port-scanner-in-go-d436e48fde87

package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			//Port is not open. Ignore
			continue
		}
		conn.Close()
		fmt.Printf("Port %d is open\n", i)
	}
}
