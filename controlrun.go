package controlrun

import (
	"errors"
	"net/http"
)

var (
	firehost string = "" // Firebase Realtime database
	token    string = "" // Project/token access
)

// Setup the access endpoint
func Set(firebase_server string, project_token string) {
	firehost = firebase_server
	token = project_token
}

// Check if have access
func Run() (bool, error) {
	if len(firehost) == 0 || len(token) == 0 {
		return false, errors.New("Setup missing!")
	}
	resp, err := http.Get("https://" + firehost + ".firebaseio.com/" + token + ".json")
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	return false, nil
}
