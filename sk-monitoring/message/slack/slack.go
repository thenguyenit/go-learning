package slack

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

//Slack type is a Slack information to connect to Slack channel
type Slack struct {
	APIEndpoint string
	Token       string
	Channel     string
	UserName    string
}

type Query map[string]string

//New will inital Slack configuration
func New() Slack {
	var slack Slack
	slack.APIEndpoint = os.Getenv("SLACK_API_END_POINT")
	slack.Token = os.Getenv("SLACK_TOKEN")
	slack.Channel = os.Getenv("SLACK_MONITORING_CHANNEL")
	slack.UserName = os.Getenv("SLACK_USERNAME")

	return slack
}

//ChatPostMessage will send a message to a Slack Channel
func (slack Slack) ChatPostMessage(message string) {
	slackFunctionName := "chat.postMessage"
	endPoint := slack.APIEndpoint + slackFunctionName

	v := url.Values{}
	v.Set("token", slack.Token)
	v.Set("channel", slack.Channel)
	v.Set("text", message)
	v.Set("username", slack.UserName)

	req, err := http.NewRequest("POST", endPoint+"?"+v.Encode(), nil)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	bs := make([]byte, 99999)
	response.Body.Read(bs)
	fmt.Println(string(bs))
}
