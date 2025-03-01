package handlers

import ( 
        "fmt"
        )
 

func CheckIPAddress(ipaddrs string) bool {
        if ipaddrs == "127.0.0.1" { 
                fmt.Println("its inside table xd")
                return true
        }
        return false
}
