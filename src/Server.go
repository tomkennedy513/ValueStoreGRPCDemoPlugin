package src

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pivotal/test/proto"
)

type Server struct {
}

func (s *Server) SetValue(ctx context.Context, in *proto.SetValueRequest) (*proto.SetValueResponse, error) {
	fmt.Println(in)
	secret := SetValue(in.Path, in.Value)
	if secret == nil {
		return &proto.SetValueResponse{}, nil
	}
	fmt.Println(secret)
	createdAt := secret.Data["created_time"].(string)
	version, _ := secret.Data["version"].(json.Number).Int64()

	id := EncodeId(in.Path)

	return &proto.SetValueResponse{
		Id:      id,
		Path:      in.Path,
		CreatedAt: createdAt,
		Version:   version,
	}, nil
}

func (s *Server) GetValueById(ctx context.Context, in *proto.GetValueByIdRequest) (*proto.GetValueResponse, error) {
	fmt.Println(in)

	path := DecodeId(in.Id)
	secret := GetValue(path)
	if secret == nil {
		return &proto.GetValueResponse{}, errors.New("value does not exist")
	}
	fmt.Println(secret.Data)
	data := secret.Data["data"].(map[string]interface{})

	return &proto.GetValueResponse{
		Value: data[path].(string),
	}, nil
}

func (s *Server) GetValueByPath(ctx context.Context, in *proto.GetValueByPathRequest) (*proto.GetValueResponse, error) {
	fmt.Println(in)

	secret := GetValue(in.Path)
	if secret == nil {
		return &proto.GetValueResponse{}, errors.New("value does not exist")
	}
	fmt.Println(secret.Data)
	data := secret.Data["data"].(map[string]interface{})

	return &proto.GetValueResponse{
		Value: data[in.Path].(string),
	}, nil
}

func (s *Server) DeleteValueById(ctx context.Context, in *proto.DeleteValueByIdRequest) (*proto.DeleteValueResponse, error) {
	fmt.Println(in)

	path := DecodeId(in.Id)

	_ = DeleteValue(path)


	return &proto.DeleteValueResponse{
		Id: in.Id,
		Path: path,
	}, nil
}

func (s *Server) DeleteValueByPath(ctx context.Context, in *proto.DeleteValueByPathRequest) (*proto.DeleteValueResponse, error) {
	fmt.Println(in)

	id := EncodeId(in.Path)

	_ = DeleteValue(in.Path)

	return &proto.DeleteValueResponse{
		Id: id,
		Path: in.Path,
	}, nil
}


