package sender

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	accountSID   string
	authToken    string
	fromNumber   string
	httpClient   = &http.Client{}
	ErrMsgFailed = errors.New("")
)

func init() {
	accountSID = os.Getenv("ACCOUNT_SID")
	authToken = os.Getenv("AUTH_TOKEN")
}

func sendSMS(to, body string) error {
	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", fromNumber)
	msgData.Set("Body", body)
	msgDataReader := strings.NewReader(msgData.Encode())

	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSID + "/Messages.json"
	req, err := http.NewRequest(http.MethodPost, urlStr, msgDataReader)
	if err != nil {
		return err
	}

	// Set the auth header.
	req.SetBasicAuth(accountSID, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	// TODO: process response and detect errors
	_ = res
	return nil
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func ProcessMessage(ctx context.Context, m PubSubMessage) error {

	return nil
}
