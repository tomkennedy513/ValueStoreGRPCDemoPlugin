package main

import (
	"context"
	"encoding/json"
	"fmt"
)

type Server struct {
}

func (s *Server) StoreValue(ctx context.Context, in *StoreValueRequest) (*StoreValueResponse, error){
	fmt.Println(in)
	secret := StoreValue(in)
	fmt.Println(secret)
	//data := secret.Data["data"].(map[string]interface{})
	createdAt := secret.Data["created_time"].(string)
	version := secret.Data["version"].(json.Number)
	print(version)

	return &StoreValueResponse{
		Path: "here",
		CreatedAt: createdAt,
		Version: 1,
	}, nil
}


func (s *Server) GetValue(ctx context.Context, in *GetValueRequest) (*GetValueResponse, error){
	fmt.Println(in)
	secret := GetValue(in)
	fmt.Println(secret.Data)
	data := secret.Data["data"].(map[string]interface{})
	return &GetValueResponse{
		Value: data[in.Path].(string),
	}, nil
}


