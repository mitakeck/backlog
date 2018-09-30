package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/k0kubun/pp"
)

var (
	config Config
)

func callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query()["code"]
	fmt.Println(code)

	io.WriteString(w, "Success to get token.  Plase close this tab.\n")

	go func() {
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}()
}

func listen() error {
	http.HandleFunc("/callback", callback)

	if err := http.ListenAndServe(":5432", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return nil
}

func auth(space string) error {
	base, err := url.Parse(fmt.Sprintf("https://%s.backlog.jp", space))
	if err != nil {
		return err
	}
	endPoint := "/OAuth2AccessRequest.action"

	params := url.Values{}
	params.Set("response_type", "code")
	params.Set("client_id", ClientID)
	url := ComposeURL(*base, endPoint, params)

	openBrowser(url)

	return nil
}

// Login : login to backlog
func Login(space string) error {
	err := auth(space)
	if err != nil {
		return err
	}

	listen()

	return nil
}

// func AccessToken(space string) (Config, error) {
//
// }

// Refresh : refresh access token
func (c *Client) Refresh() error {
	params := url.Values{}
	params.Add("grant_type", "refresh_token")
	params.Add("client_id", ClientID)
	params.Add("client_secret", ClientSecret)
	params.Add("refresh_token", c.APIRefreshToken)

	body, err := c.execute("POST", "/api/v2/oauth2/token", params, false)
	if err != nil {
		return err
	}

	var model AccessToken
	json.Unmarshal(body, &model)

	c.APIAccessToken = model.AccessToken
	pp.Println(model)
	return nil
}

// Issue : get issue detail
func (c *Client) Issue(ticketID string) (Issue, error) {
	endPoint := "/api/v2/issues/" + ticketID

	body, err := c.execute("GET", endPoint, url.Values{}, true)
	if err != nil {
		return Issue{}, err
	}

	var model Issue
	json.Unmarshal(body, &model)

	return model, nil
}
