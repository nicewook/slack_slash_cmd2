package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/nicewook/slack_slash_cmd2/slack/slash"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome, I'm Timebot\nI can convert KST <-> PST/PDT")
}

//var slackSigningToken string

func mustLookupEnv(env string) string {
	ret, ok := os.LookupEnv(env)

	if !ok {
		log.Fatalf("Environment Variable %v is not available!\n", env)
	}
	return ret
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slack.slackSigningToken = mustLookupEnv("SLACK_SIGNING_SECRET")

	r := httprouter.New()
	r.GET("/", index)
	r.POST("/slash", slash.Handler)

	fmt.Println("[INFO] Server listening")
	log.Fatal(http.ListenAndServe("", r))

}
