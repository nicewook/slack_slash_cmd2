package slack2

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// SlackSigningToken get environment variable value
var SlackSigningToken string

// Handler deals with slash commands
// Check HTTP Request and ParseForm()
func Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// verify token - compare received r's TOKEN and server' slackSigning Token
	if ok := VerifyRequest(r, []byte(SlackSigningToken)); ok == false {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("VerifyRequest()")
		fmt.Println("VerifyRequest()")
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("r.ParseForm()")
		fmt.Println("r.ParseForm()")
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
