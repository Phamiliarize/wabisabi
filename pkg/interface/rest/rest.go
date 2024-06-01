// Package rest implements the HTTP interface for wabisabi

// In other words, if a resource is to be exposed over HTTP, it's HTTP
// interface layer concerns should be handled in this package
package rest

import (
	"net/http"

	"github.com/Phamiliarize/wabisabi/pkg/application"
)

type restAPI struct {
	application application.Application
}

func NewRestInterface(application application.Application) *restAPI {
	return &restAPI{
		application: application,
	}
}

func (r *restAPI) PostCreateSession(w http.ResponseWriter, req *http.Request) {
	var reqBody CreateSessionJSON

	err := jsonRequest(req, &reqBody)
	if err != nil {
		switch err.Error() {
		case "invalid_content_type":
			jsonResponse(
				http.StatusBadRequest,
				GenericErrorMessageJSON{Message: "Expected Content-Type of 'application/json'"},
				w,
			)
		default:
			jsonResponse(http.StatusInternalServerError, nil, w)
		}
		return
	}

	_, err = r.application.CreateSession(application.Session(reqBody))
	if err != nil {
		switch v := err.(type) {
		case application.VadliationErrors:
			jsonResponse(http.StatusBadRequest, v, w)

		default:
			jsonResponse(http.StatusInternalServerError, nil, w)
		}
		return
	}

	response := map[string]interface{}{"status": "ok"}
	jsonResponse(http.StatusOK, response, w)
}

func (r *restAPI) PostValidateSession(w http.ResponseWriter, req *http.Request) {
	var reqBody ValidateSessionJSON

	err := jsonRequest(req, &reqBody)
	if err != nil {
		switch err.Error() {
		case "invalid_content_type":
			jsonResponse(
				http.StatusBadRequest,
				GenericErrorMessageJSON{Message: "Expected Content-Type of 'application/json'"},
				w,
			)
		default:
			jsonResponse(http.StatusInternalServerError, nil, w)
		}
		return
	}

	token := application.Token(reqBody)

	_, err = r.application.ValidateSession(token)
	if err != nil {
		switch v := err.(type) {
		case application.VadliationErrors:
			jsonResponse(http.StatusBadRequest, v, w)

		default:
			jsonResponse(http.StatusInternalServerError, nil, w)
		}
		return
	}

	response := map[string]interface{}{"status": "ok"}
	jsonResponse(http.StatusOK, response, w)
}

func (r *restAPI) PostDeleteSessionByToken(w http.ResponseWriter, req *http.Request) {
	var reqBody DeleteSessionByTokenJSON

	err := jsonRequest(req, &reqBody)
	if err != nil {
		switch err.Error() {
		case "invalid_content_type":
			jsonResponse(
				http.StatusBadRequest,
				GenericErrorMessageJSON{Message: "Expected Content-Type of 'application/json'"},
				w,
			)
		default:
			jsonResponse(http.StatusInternalServerError, nil, w)
		}
		return
	}

	token := application.Token(reqBody)

	err = r.application.DeleteSessionByTokenId(token)
	if err != nil {
		switch v := err.(type) {
		case application.VadliationErrors:
			jsonResponse(http.StatusBadRequest, v, w)

		default:
			jsonResponse(http.StatusInternalServerError, nil, w)
		}
		return
	}

	response := map[string]interface{}{"status": "ok"}
	jsonResponse(http.StatusOK, response, w)
}

func (r *restAPI) PostDeleteSessionByUser(w http.ResponseWriter, req *http.Request) {
	var reqBody DeleteSessionByUserJSON

	err := jsonRequest(req, &reqBody)
	if err != nil {
		switch err.Error() {
		case "invalid_content_type":
			jsonResponse(
				http.StatusBadRequest,
				GenericErrorMessageJSON{Message: "Expected Content-Type of 'application/json'"},
				w,
			)
		default:
			jsonResponse(http.StatusInternalServerError, nil, w)
		}
		return
	}

	user := application.User(reqBody)

	err = r.application.DeleteSessionByUserId(user)
	if err != nil {
		switch v := err.(type) {
		case application.VadliationErrors:
			jsonResponse(http.StatusBadRequest, v, w)

		default:
			jsonResponse(http.StatusInternalServerError, nil, w)
		}
		return
	}

	response := map[string]interface{}{"status": "ok"}
	jsonResponse(http.StatusOK, response, w)
}
