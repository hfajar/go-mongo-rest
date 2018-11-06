package services

import "gopkg.in/mgo.v2"

// Connect is for connection database
func Connect() (*mgo.Session, error) {
	var session, err = mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	return session, nil
}
