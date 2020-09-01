package main

import (
	"context"
	"fmt"
	"github.com/dyweb/gommon/util/httputil"
	"github.com/google/go-github/v29/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"os"
)

func main() {

	var token = os.Getenv("GITHUB_TOKEN")

	hc := httputil.NewUnPooledClient()
	if token != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: token,
		})
		hc = oauth2.NewClient(context.Background(), ts)
	}

	client := github.NewClient(hc)

	title := "测试自动开issues"
	assignee := []string{"soyum2222"}

	body := "测试自动开issues:"
	req := &github.IssueRequest{
		Title: &title,
		Labels: &[]string{
			"bug",
		},
		Assignees: &assignee,
		Body:      &body,
	}

	issue, resp, err := client.Issues.Create(context.Background(), "soyum2222", "actions", req)
	if err != nil {
		fmt.Println(err)
	}
	body1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body1))
	fmt.Println("Created new issue %d %s", issue.GetNumber(), issue.GetTitle())

}
