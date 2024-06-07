//go run concurrent_brute_force.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"net"
	"strings"
	"time"
)


func main() {
    //read wordlist from file
    filePath := "wordlist.txt"
    readFile, err := os.Open(filePath)

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string

    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }

    readFile.Close()

    //read common azure resource from file
    azPath := "azure.txt"
    readazFile, err := os.Open(azPath)

    if err != nil {
        fmt.Println(err)
    }
    azScanner := bufio.NewScanner(readazFile)
    azScanner.Split(bufio.ScanLines)
    var azLines []string

    for azScanner.Scan() {
        azLines = append(azLines, azScanner.Text())
    }

    readazFile.Close()


    //loop thorugh both lists
    for _, azline := range azLines {
        for _, line := range fileLines {
            //fmt.Println(line+azline)
		go dostuff(line,azline)
		}
	time.Sleep(1 * time.Second)
	fmt.Println(azline," Search Complete")
	}

	var input string
	fmt.Scanln(&input)

}

func dostuff(n string, d string) {
	time.Sleep(1 * time.Second)

        base := "evilcorp"

	//base query
	dnsquery := base+n+d
	ips, err_ := net.LookupIP(dnsquery)
	if err_ != nil {
		if strings.Contains(err_.Error(),"no such host") {
			//fmt.Fprintf(os.Stderr, "Invalid Target: %v\n", err_)
		} else {
			fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err_)
		}
	} else {
	fmt.Printf(dnsquery+" IN A %s\n", ips[0].String())
	}

        // reversed base query
	revdnsquery := n+base+d
	rips, err_ := net.LookupIP(revdnsquery)
	if err_ != nil {
		if strings.Contains(err_.Error(),"no such host") {
			//fmt.Fprintf(os.Stderr, "Invalid Target: %v\n", err_)
		} else {
			fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err_)
		}
		//os.Exit(1)
	} else {

	fmt.Printf(revdnsquery+" IN A %s\n", rips[0].String())
	}

}
