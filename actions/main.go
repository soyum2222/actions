package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dyweb/gommon/util/httputil"
	"github.com/google/go-github/v29/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"os"
)

func ReadFileToString(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

type Issue struct {
	Title    string   `json:"title"`
	Body     string   `json:"body"`
	Labels   []string `json:"labels"`
	Assignee []string `json:"assignee"`
}

func main() {


	var cfg string

	flag.StringVar(&cfg,"cfg","","")


	data ,err:=ioutil.ReadFile(cfg)
	if err!=nil{
		panic(err)
	}


	issueInfo := Issue{}
	err = json.Unmarshal(data,&issueInfo)
	if err!=nil {
		panic(err)
	}


	var token = os.Getenv("GITHUB_TOKEN")

	hc := httputil.NewUnPooledClient()
	if token != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: token,
		})
		hc = oauth2.NewClient(context.Background(), ts)
	}

	client := github.NewClient(hc)

	req := &github.IssueRequest{
		Title:     &issueInfo.Title,
		Labels:    &issueInfo.Labels,
		Assignees: &issueInfo.Assignee,
		Body:      &issueInfo.Body,
	}

	issue, _, err := client.Issues.Create(context.Background(), "soyum2222", "actions", req)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("Created new issue %d %s", issue.GetNumber(), issue.GetTitle())

}
