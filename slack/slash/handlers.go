package slack

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nicewook/slack_slash_cmd2/slack"
)

var slackSigningToken string

// Handler deals with slash commands
// Check HTTP Request and ParseForm()
func Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// verify token - compare received r's TOKEN and server' slackSigning Token
	if ok := slack.VerifyRequest(r, []byte(slackSigningToken)); ok == false {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// print all POST values
	for k, v := range r.PostForm {
		fmt.Printf("%v: %v\n", k, v)
	}

	response := fmt.Sprintf("Request accepted")
	w.Write([]byte(response))
}

// Handler deals with slash commands
// func Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	cmd := r.PostFormValue("command")

// 	// TODO: validateToken

// 	switch cmd {
// 	case "/time":
// 		response := fmt.Sprintf("You requested for KST <-> PST/PDT for %+v", r)
// 		w.Write([]byte(response))

// 	default:
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// }
