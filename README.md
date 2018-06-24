# ControlRun
_Control the access of your apps with Firebase_

## Setting Firebase
 1. Create a Firebase Project  https://firebase.google.com
 2. Go to Database options and choice to create a "Realtime Database".
 3. Go to rules and paste:
 ```
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
```
 4. Go to Data and import or create an struct like this code:
```
{
 "appName" : {
  "tokenCode" : 0
 }
}
```
## Start using it
 1.  Download and install it:
 ```
    $ go get github.com/jonathanhecl/controlrun
 ``` 
 2.  Import it in your code:
 ```
    import "github.com/jonathanhecl/controlrun"
 ``` 
 3. Setup your Firebase:
 ```
    controlrun.Set("[ProjectName]", "[appName]/[tokenCode]")
 ``` 
 4. Control the access:
 ```
    access, _ := controlrun.Run()
    if !access {
    	panic("Permission denied")
    }
 ```
