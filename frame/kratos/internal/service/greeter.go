package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"github.com/wangle201210/studyGo/tools/bilibili/service"

	pb "kratos/api/helloworld/v1"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer
	s *service.Service
}

func NewGreeterService() *GreeterService {
	s := service.New()
	return &GreeterService{
		s: s,
	}
}

func (s *GreeterService) List(ctx context.Context, req *pb.ListRequest) (resp *pb.ListReply, err error) {
	info, err := s.s.GetUserListInfo(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		log.Error(ctx, "GetUserListInfo err %+v", err)
		return
	}
	log.Info(ctx, "GetUserListInfo %+v", info)
	var list []*pb.Userinfo
	for _, userInfo := range info {
		item := new(pb.Userinfo)
		copier.Copy(item, userInfo)
		list = append(list, item)
	}
	return &pb.ListReply{
		List: list,
	}, nil
}
