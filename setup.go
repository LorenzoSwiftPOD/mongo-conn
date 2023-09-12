package mongoconn

import (
	"os"
)

var Mode string = "production"

type tokens struct {
	uri           string
	user          string
	port          string
	host          string
	password      string
	authDatabase  string
	database      string
	collection    string
	authMechanism string
}

func (t *tokens) init() {
	t.uri = os.Getenv("MongoURI")
	t.port = os.Getenv("MongoPort")
	t.user = os.Getenv("MongoUser")
	t.password = os.Getenv("MongoPassword")
	t.authDatabase = os.Getenv("MongoAuthDatabase")
	t.authMechanism = os.Getenv("MongoAuthMechanism")
	if Mode == "local" {
		t.uri = os.Getenv("LocalMongoURI")
		t.port = os.Getenv("LocalMongoPort")
		t.user = os.Getenv("LocalMongoUser")
		t.password = os.Getenv("LocalMongoPassword")
		t.authDatabase = os.Getenv("LocalMongoAuthDatabase")
		t.authMechanism = os.Getenv("LocalMongoAuthMechanism")
	}

}
