package main

import (
	"context"
	"fmt"
)

type Server struct {
}

func (s *Server) StoreValue(ctx context.Context, in *StoreValueRequest) (*ValueResponse, error){
	fmt.Println(in)
	secret := StoreValue(in)
	fmt.Println(secret)

	return &ValueResponse{}, nil
}


func (s *Server) GetValue(ctx context.Context, in *GetValueRequest) (*ValueResponse, error){
	fmt.Println(in)
	secret := GetValue(in)
	fmt.Println(secret)
	return &ValueResponse{}, nil
}


