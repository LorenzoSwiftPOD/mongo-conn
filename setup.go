package mongoconn

import (
	"os"
)

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
	t.port = os.Getenv("SwiftMongoPort")
	t.user = os.Getenv("SwiftMongoUser")
	t.password = os.Getenv("SwiftMongoPassword")
	t.authDatabase = os.Getenv("SwiftMongoAuthDatabase")
	t.authMechanism = os.Getenv("SwiftMongoAuthMechanism")

}
