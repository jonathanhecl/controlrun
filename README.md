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
The tokenCode can be what you want, like the client name.

 5. If you want to activate an tokenCode, only need put 1 in the value.

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
    err := controlrun.Run()
    if err != nil {
    	panic("Permission denied")
    }
 ```

## How use hash verification
 1.  Run in a local file or an any test:
```
    hash := controlrun.Set("[ProjectName]", "[appName]/[tokenCode]")
    println(hash)
 ``` 
2.   Copy your hash generated, and use it in the Run() function:
 ```
    err := controlrun.Run("da39a3ee5e6b4b0d3255bfef95601890afd80709")
    if err != nil {
    	panic("Permission denied")
    }
 ```
