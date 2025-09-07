package main

type authenticationInfo struct {
	username string
	password string
}

// create the method below

func (auth authenticationInfo) getBasicAuth() string {
	s := "Authorization: Basic " + auth.username + ":" + auth.password
	return s
}
