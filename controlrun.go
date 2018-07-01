package controlrun

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"net/http"
)

var (
	firehost string = "" // Firebase Realtime database
	token    string = "" // Project/token access
)

// Setup the access endpoint
func Set(firebase_server string, project_token string) string {

	_hash := sha1.New()

	firehost = firebase_server
	token = project_token

	_hash.Write([]byte(firehost + token))
	_sha := hex.EncodeToString(_hash.Sum(nil))

	return _sha

}

// Check if have access
func Run(hash ...string) error {

	if len(firehost) == 0 || len(token) == 0 {
		return errors.New("Setup missing!")
	}

	if len(hash) > 0 {
		_hash := sha1.New()
		_hash.Write([]byte(firehost + token))
		_sha := hex.EncodeToString(_hash.Sum(nil))

		if _sha != hash[0] {
			return errors.New("Permission denied!")
		}
	}

	resp, err := http.Get("https://" + firehost + ".firebaseio.com/" + token + ".json")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return errors.New("Permission denied!")

}
