package main

import (
	"fmt"
	"net"
	"time"
	"flag"
	"strings"
	"errors"
)

func checkSvc(repeats int, destination string, success chan bool) {
	for i := 0; i < repeats; i++ {
		err := checkPort(destination)
		if err != nil {
			fmt.Println(err)
			time.Sleep(5 * time.Second)
		} else {
			fmt.Println("Successfully connected to " + destination)
			success <- true
			return
		}
	}
	success <- false

}

func checkPort(destination string) error {
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		return err
	}
	conn.SetDeadline(time.Now().Add(1 * time.Second))
	return err

}
func parseFlags() (int, []string, error) {
	urlString := flag.String("services", "", "Services dns:port")
	repeats := flag.Int("repeats", 10, "Dead repeats")
	flag.Parse()
	urls := strings.Split(*urlString, ",")
	if len(urls) >= 0 {
		return *repeats, urls, nil
	}
	return *repeats, nil, errors.New("not enough arguments")



}

func main() {
	repeats, urls, err := parseFlags()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, url :=range urls {
		fmt.Println(url)
	}
	success := make(chan bool)
	for _, url := range urls {
		go checkSvc(repeats, url, success)
	}
	for i:=1;i <=len(urls); i++ {
		result := <- success
		fmt.Println(result)
		if result == false {
			return
		}
	}


}
