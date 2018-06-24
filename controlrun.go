// * Step 1: Create a Firebase Project ( https://firebase.google.com )
//
// * Step 2: Go to https://console.firebase.google.com/project/[ProjectName]/database
//
// * Step 3: Choice "Realtime Database"
//
// * Step 4: Go to rules an paste:
//
/*
	{
	  "rules": {
		".read": false,
	    ".write": false,
		"$project": {
	    	".read": false,
			"$token": {
	        	".read": "data.val() == 1"
	      	}
	    }
	  }
	}
*/
// * Step 5: Go to Data and import or create an struct like this code:
//
/*
	{
	  "appName" : {
	    "tokenCode" : 0
	  }
	}
*/
// * Step 6: Add in your code:
/*
	import (
		"github.com/jonathanhecl/controlrun"
	)
*/
// * Step 7: Setup it:
/*
	Set("[ProjectName]", "[appName]/[tokenCode]")
*/
// * Step 8: Use it:
/*
	access, _ := Run()
	if !access {
		panic("Permission denied")
	}
*/
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
