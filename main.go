package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {

	fmt.Println("Enter Message: ")
	inputReader := bufio.NewReader(os.Stdin)
	message, _ := inputReader.ReadString('\n')
	fmt.Println("Enter webhook: ")
	var second string
	fmt.Scanln(&second)
	fmt.Println("Enter name: ")
	Wname, _ := inputReader.ReadString('\n')
	fmt.Println("Enter avatar url: ")
	avurl, _ := inputReader.ReadString('\n')
	data := url.Values{
		"content":    {message},
		"username":   {Wname},
		"avatar_url": {avurl},
	}
	for {
		resp, err := http.PostForm(second, data)

		if err != nil {
			log.Fatal(err)
		}

		var res map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&res)

		if resp.StatusCode >= 200 && resp.StatusCode <= 204 {
			fmt.Println(message, "sent!")
		} else if resp.StatusCode >= 429 {
			fmt.Println("Rate limited")
		} else if resp.StatusCode >= 404 {
			fmt.Println("Webhook Deleted")
			time.Sleep(8 * time.Second)
			break
		}

	}
}
