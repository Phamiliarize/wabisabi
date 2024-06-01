/*
Serves the wabisabi application over an HTTP interface

Usage:

	server [flags]

The flags are:

	-p
		Set a specific port to run the server on: expects a unsigned integer

	-u
		Takes the string regex of how to validate your user ID; defaults to UUID.

It is essentially a simple HTTP Server with little magic.
*/
package main

import (
	"flag"
	"fmt"
	"net/http"
	"regexp"

	"github.com/Phamiliarize/wabisabi/pkg/adapter/sqlite"
	"github.com/Phamiliarize/wabisabi/pkg/application"
	"github.com/Phamiliarize/wabisabi/pkg/interface/rest"
)

const UUID_REGEXP = "[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}"

func main() {
	var (
		p = flag.Uint("p", 3000, "Port to server from (Default: 3000)")
		u = flag.String("u", UUID_REGEXP, "Regex string for validating User ID format (Default: UUID)")
	)
	flag.Parse()

	userIdRegExp := regexp.MustCompile(*u)

	datastore := sqlite.NewSqLite()

	wabisabi := application.NewWabisabi(datastore, userIdRegExp)

	restAPI := rest.NewRestInterface(wabisabi)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/session/create", restAPI.PostCreateSession)
	mux.HandleFunc("POST /api/session/validate", restAPI.PostValidateSession)
	mux.HandleFunc("POST /api/session/delete/token", restAPI.PostDeleteSessionByToken)
	mux.HandleFunc("POST /api/session/delete/user", restAPI.PostDeleteSessionByUser)

	http.ListenAndServe(fmt.Sprintf(":%d", *p), mux)
}
