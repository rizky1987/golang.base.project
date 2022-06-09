package mongo

import (
	"time"

	"gopkg.in/mgo.v2"
)

type Info struct {
	Hostname string
	Database string
	Username string
	Password string
}

func Connect(host, database, user, password, port string) (*mgo.Session, error) {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:     []string{host},
		Timeout:   60 * time.Second,
		Database:  database,
		Username:  user,
		Password:  password,
		Source:    "admin",
		Mechanism: "SCRAM-SHA-1",
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session, err
}
