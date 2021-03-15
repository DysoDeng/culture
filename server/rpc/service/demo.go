package service

import (
	"context"
	"culture/cloud/base/server/rpc/proto"
	"errors"
	"log"
	"time"
)

type DemoService struct {}

func (s *DemoService) UserInfo(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	log.Println("DemoService UserInfo " + time.Now().Format("2006-01-02 15:04:05"))
	if req.Uid > 0 {
		return &proto.Response{
			Id:       req.Uid,
			Username: "dds",
			Nickname: "栖枝",
		}, nil
	} else {
		return nil, errors.New("id 不能小于1")
	}
}
