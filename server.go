package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/nicewook/slack_slash_cmd2/slack2"
	// "github.com/nicewook/slack_slash_cmd2/slack"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome, I'm Timebot\nI can convert KST <-> PST/PDT")
}

//var SlackSigningToken string

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

	sst := mustLookupEnv("SLACK_SIGNING_SECRET")
	slack2.SlackSigningToken = sst

	r := httprouter.New()
	r.GET("/", index)
	r.POST("/slash", slack2.Handler)

	fmt.Println("[INFO] Server listening")
	log.Fatal(http.ListenAndServe("", r))

}
