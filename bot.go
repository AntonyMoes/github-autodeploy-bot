package main

import (
	"bytes"
	"fmt"
	"github.com/google/go-github/github"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

var (
	config *Config
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

	switch event.(type) {
	case *github.PushEvent:
		ePush := event.(*github.PushEvent)

		repoName := *ePush.Repo.FullName
		repo := config.Repos[repoName]
		fmt.Printf("Repo: %s\n", repo)

		refSlice := strings.Split(*ePush.Ref, "/")
		branch := refSlice[len(refSlice)-1]
		fmt.Printf("Branch: %s\n", branch)

		//cmd := exec.Command("./front_layout_update.sh")
		fmt.Printf("Cmd: %s\n", repo[branch])
		cmd := exec.Command(repo[branch])
		cmd.Stdin = strings.NewReader("")

		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Output: %v \n\n", out.String())

	case *github.PullRequestEvent:
		// this is a pull request, do something with it

	default:
		log.Printf("unknown event type %s\n", github.WebHookType(r))
		return
	}
}

func main() {
	config = &Config{}

	configBytes, err := ioutil.ReadFile("webhooks.json")
	if err != nil {
		log.Fatalf("Readn't: %v", err)
	}

	err = config.UnmarshalJSON(configBytes)
	if err != nil {
		log.Fatalf("Unmarshalln't: %v", err)
	}

	log.Printf("server started on port %s\n", config.Port)
	http.HandleFunc(config.Url, handleWebhook)
	log.Fatal(http.ListenAndServe(config.Port, nil))
}
