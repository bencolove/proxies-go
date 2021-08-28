package main

import (
	"fmt"
	"log"

	"com.roger.httproxy.go/lib/httproxies"
	"com.roger.httproxy.go/lib/scrape"
)

func main() {
	// url := "http://193.149.225.9"
	// if ok, err := checkProxyHealth(url); err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	if ok {
	// 		fmt.Println("proxy OK")
	// 	} else {
	// 		fmt.Println("proxy DOWN")
	// 	}
	// }

	err := scrape.CollectProxy()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Done")
	}
}

func checkProxyHealth(proxyURL string) (bool, error) {
	if client, err := httproxies.ProxiedClient(proxyURL); err != nil {
		return false, err
	} else {
		if resp, err := client.Get("http://httpbin.org/get"); err != nil {
			return false, err
		} else if resp.StatusCode != 200 {
			return false, nil
		} else {
			return true, nil
		}
	}
}
