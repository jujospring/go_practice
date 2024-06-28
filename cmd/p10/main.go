package main

type authenticationInfo struct {
	username string
	password string
}

// create the method below
func (auth authenticationInfo) getBasicAuth() string {
	return "Authorization: Basic " + auth.username + ":" + auth.password
}

// var auth = authenticationInfo{
// 	username: "bug",
// 	password: "jones",
// }
