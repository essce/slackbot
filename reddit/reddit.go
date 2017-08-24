package reddit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/standielpls/slackbot"
)

type RedditProcess struct {
	URL string
}

// GetRedditPosts gets reddit posts.
func (r *RedditProcess) GetRedditPosts() string {
	url := fmt.Sprintf(r.URL)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "slacka/1.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// bodie, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("error reading from resp body")
	// }

	//fmt.Println("reddit:" + string(bodie))
	var redditResp slackbot.Reddit
	err = json.NewDecoder(resp.Body).Decode(&redditResp)
	if err != nil {
		return fmt.Sprintf("error decoding json: %v", err)
	}
	var buffer bytes.Buffer
	for i := 0; i < 5; i++ {
		buffer.WriteString(redditResp.Data.Children[i].Data.URL + "\n")
	}

	return fmt.Sprintf(buffer.String())
	//return ""
}
