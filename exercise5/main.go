package main

import (
	"./ipaddress"
	"fmt"
)

func main() {
	ip := ipaddress.GetIpAddress("25525510124")
	fmt.Printf("All possible valid IP addresses: %v\n", ip)
}
