package main

import (
	"bytes"
	"strings"
	"fmt"
	"os/exec"
	"log"
	"net/http"
	"github.com/google/go-github/github"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := github.ValidatePayload(r, []byte("nya"))
	if err != nil {
		log.Printf("error validating request body: err=%s\n", err)
		return
	}
	defer r.Body.Close()


	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("could not parse webhook: err=%s\n", err)
		return
	}

	switch e := event.(type) {
	case *github.PushEvent:
		// this is a commit push, do something with it
		fmt.Println("Got a push event")
		cmd := exec.Command("./update.sh")
		cmd.Stdin = strings.NewReader("")
		fmt.Println("\nBUILDING READER")
		var out bytes.Buffer
		cmd.Stdout = &out
		fmt.Println("BINDING OUTPUT STREAM")
		err := cmd.Run()
		fmt.Println("COMMAND EXECUTED")
		if err != nil {
			log.Fatal(err)
		}
	fmt.Printf("Output: %v \n\n",out.String())
	case *github.PullRequestEvent:
		// this is a pull request, do something with it
	case *github.WatchEvent:
		// https://developer.github.com/v3/activity/events/types/#watchevent
		// someone starred our repository
		if e.Action != nil && *e.Action == "starred" {
			fmt.Printf("%s starred repository %s\n",
				*e.Sender.Login, *e.Repo.FullName)
		}
	default:
		log.Printf("unknown event type %s\n", github.WebHookType(r))
		return
	}
}

func main() {
	log.Println("server started")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
