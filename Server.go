package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
)

type Server struct {
}

type UuidData struct {
	Path string
	Version int64
	CreatedAt string
}

func (s *Server) StoreValue(ctx context.Context, in *StoreValueRequest) (*StoreValueResponse, error){
	fmt.Println(in)
	secret := StoreValue(in)
	if secret == nil {
		return &StoreValueResponse{}, nil
	}
	fmt.Println(secret)
	createdAt := secret.Data["created_time"].(string)
	version, _ := secret.Data["version"].(json.Number).Int64()

	uuid := encodeUuid(&UuidData{
		Path:in.Path,
		Version:version,
		CreatedAt:createdAt,
	})

	return &StoreValueResponse{
		Uuid: uuid,
		Path: in.Path,
		CreatedAt: createdAt,
		Version: version,
	}, nil
}


func (s *Server) GetValueByPath(ctx context.Context, in *GetValueByPathRequest) (*GetValueResponse, error){
	fmt.Println(in)
	secret := GetValue(in.Path)
	if secret == nil {
		return &GetValueResponse{}, nil
	}
	fmt.Println(secret.Data)
	data := secret.Data["data"].(map[string]interface{})
	return &GetValueResponse{
		Value: data[in.Path].(string),
	}, nil
}
func (s *Server) GetValueByUUID(ctx context.Context, in *GetValueByUUIDRequest) (*GetValueResponse, error){
	fmt.Println(in)

	path := decodeUuid(in.Uuid).Path

	secret := GetValue(path)
	if secret == nil {
		return &GetValueResponse{}, nil
	}
	fmt.Println(secret.Data)
	data := secret.Data["data"].(map[string]interface{})
	return &GetValueResponse{
		Value: data[path].(string),
	}, nil
}

func decodeUuid(uuid string) *UuidData {
	decodedBytes, err := hex.DecodeString(uuid)
	if err != nil {
		log.Fatal(err)
	}

	buffer := bytes.NewBuffer(decodedBytes)
	var data UuidData
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	return &data

}

func encodeUuid(data *UuidData) string {
	var dataBuffer bytes.Buffer
	enc := gob.NewEncoder(&dataBuffer)

	err := enc.Encode(data)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(dataBuffer.Bytes())
}
