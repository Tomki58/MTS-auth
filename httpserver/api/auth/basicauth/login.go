package basicauth

import (
	"bytes"
	"os"

	"github.com/golang-jwt/jwt"
)

type BasicAuthorizator struct {
	credentials map[string]string
}

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func New(pathToConfig string) (*BasicAuthorizator, error) {
	data, err := os.ReadFile(pathToConfig)
	if err != nil {
		return nil, err
	}

	creds := make(map[string]string)
	for _, auth := range bytes.Split(data, []byte{'\n'}) {
		cs := bytes.IndexByte(auth, ':')
		if cs < 0 {
			continue
		}
		username, password := auth[:cs], auth[cs+1:]
		creds[string(username)] = string(password)
	}

	return &BasicAuthorizator{
		credentials: creds,
	}, nil
}

func (b *BasicAuthorizator) CheckCredentials(username, password string) bool {
	if _, ok := b.credentials[username]; ok {
		return true
	}

	return false
}
