package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
)

func StoreValue(request *StoreValueRequest) *api.Secret {

	client := getClient()

	payload := map[string]interface{}{
		"data": map[string]interface{}{
			request.Path: request.Value,
		},
		"options": map[string]interface{}{},
	}

	res, err := client.Write(fmt.Sprintf("/secret/data/%s", request.Path), payload)
	if err != nil {
		fmt.Println(err)
	}
	return res
}


func GetValue(path string) *api.Secret {
	client := getClient()

	secret, err := client.Read(fmt.Sprintf("/secret/data/%s", path))
	if err != nil {
		fmt.Println(err)
	}
	return secret
}


func getClient() *api.Logical {
	config := api.Config{
		Address: "http://127.0.0.1:8200",

	}
	client, _ := api.NewClient(&config)
	client.SetToken("s.7PpDWHNlxSz4cAkY3Lkn1UaN")
	return client.Logical()
}
