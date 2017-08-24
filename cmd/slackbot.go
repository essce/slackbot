/*

mybot - Illustrative Slack bot in Go

Copyright (c) 2015 RapidLoop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/standielpls/slackbot"
	"github.com/standielpls/slackbot/reddit"
	"github.com/standielpls/slackbot/slack"
)

type variables struct {
	SLACK_HEDO_KEY string
}

func main() {

	v, err := verifyEnvironmentVariables()
	if err != nil {
		fmt.Println("setup error: " + err.Error())
		return
	}

	r := reddit.RedditProcess{
		URL: "https://www.reddit.com/r/nba/hot/.json",
	}
	s := slack.Slack{
		URL:   "https://slack.com/api/rtm.start?token=%s",
		Token: v.SLACK_HEDO_KEY,
	}

	// start a websocket-based Real Time API session
	ws, id := s.SlackConnect()
	fmt.Println("hedo slackbot ready\nctrl+c to exit.")

	for {
		// read each incoming message
		m, err := s.GetMessage(ws)
		if err != nil {
			log.Fatal(err)
		}

		// see if we're mentioned
		if m.Type == "message" && strings.HasPrefix(m.Text, "<@"+id+">") {
			// if so try to parse if
			parts := strings.Fields(m.Text)
			switch parts[1] {
			case "reddit":
				go func(m slackbot.Message) {
					m.Text = r.GetRedditPosts()
					s.PostMessage(ws, m)
				}(m)
			case "stock":
				// looks good, get the quote and reply with the result
				go func(m slackbot.Message) {
					m.Text = getQuote(parts[2])
					s.PostMessage(ws, m)
				}(m)
			default:
			}
			if parts[1] == "reddit" {
				m.Text = fmt.Sprintf("ball is life\n")
				s.PostMessage(ws, m)
			}
		}
	}
}

// Get the quote via Yahoo. You should replace this method to something
// relevant to your team!
func getQuote(sym string) string {
	sym = strings.ToUpper(sym)
	url := fmt.Sprintf("http://download.finance.yahoo.com/d/quotes.csv?s=%s&f=nsl1op&e=.csv", sym)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	rows, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	if len(rows) >= 1 && len(rows[0]) == 5 {
		return fmt.Sprintf("%s (%s) is trading at $%s", rows[0][0], rows[0][1], rows[0][2])
	}
	return fmt.Sprintf("unknown response format (symbol was \"%s\")", sym)
}

func verifyEnvironmentVariables() (variables, error) {
	var v variables
	if os.Getenv("SLACK_HEDO_KEY") == "" {
		return v, fmt.Errorf("did not provide SLACK_HEDO_KEY")
	}
	v.SLACK_HEDO_KEY = os.Getenv("SLACK_HEDO_KEY")
	return v, nil
}
