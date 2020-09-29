package users

import (
	"encoding/json"
	"github.com/nasruddin/golang/elastic/clients/elasticsearch"
	"log"
)

const (
	indexUser = "users"
)

func (u *User) Save() error {
	result, err := elasticsearch.Client.Index(indexUser, u)
	if err != nil {
		log.Fatalf("Error while trying to save user in index %s ", indexUser)
	}
	u.ID = result.Id
	return nil
}

func (u *User) Get() error {
	result, err := elasticsearch.Client.Get(indexUser, u.ID)
	if err != nil {
		panic(err)
	}
	bytes, err := result.Source.MarshalJSON()
	if err := json.Unmarshal(bytes, u); err != nil {
		panic(err)
	}
	print(string(bytes))
	print(err)
	return nil
}