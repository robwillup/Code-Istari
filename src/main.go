// CodeIstari is a program to help code mentors and apprentices connect.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"	
)

func main() {
	userName := os.Args[1:][0]
	url := fmt.Sprintf("%s%s%s", "https://api.twitter.com/2/users/by?usernames=", userName, "&user.fields=description")
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TWITTER_BEARER_TOKEN")))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}