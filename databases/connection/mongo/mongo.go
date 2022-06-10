package mongo

import (
	"fmt"
	"os"
	"time"

	"example/utils"

	"gopkg.in/mgo.v2"
)

func Connect(host, database, user, password, port string) (*mgo.Session, error) {

	hostwithPort := fmt.Sprintf("%s:%s", host, port)
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:     []string{hostwithPort},
		Timeout:   60 * time.Second,
		Database:  database,
		Username:  user,
		Password:  password,
		Source:    "admin",
		Mechanism: "SCRAM-SHA-1",
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		utils.SaveErrorToApplicationInsight("failed connect to Mongo DB", "initialization_database", err.Error(), "", 0)
		os.Exit(1)
	}

	err = session.Ping()
	if err != nil {

		utils.SaveErrorToApplicationInsight("failed connect to Mongo DB", "initialization_database", err.Error(), "", 0)
		fmt.Println(err)
		os.Exit(1)
	}
	session.SetMode(mgo.Monotonic, true)
	return session, err
}
