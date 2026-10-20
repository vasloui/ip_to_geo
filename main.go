package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var api_key string

func getcreds(key string) string {
	data, err := os.ReadFile(".env")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), `\n`)

	for _, l := range lines {
		// fmt.Println(l)
		for i, _ := range l {
			// fmt.Println(l[:i])
			if string(l[:i]) == key && string(l[i]) == "=" {
				return string(l[i+1:])
			}
		}
	}
	return ""
}

func ip_to_geo(ip string) {
	url := "https://api.ipgeolocation.io/v3/ipgeo?apiKey=" + api_key + "&ip=" + ip
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func main() {
	api_key = getcreds("api_key")

	if api_key == "" {
		fmt.Println("No API key provided in .env.")
		return
	}

	argsWithProg := os.Args
	for i, s := range argsWithProg {
		if s == "-i" {
			ip_to_geo(argsWithProg[i+1])
		}
	}

}
