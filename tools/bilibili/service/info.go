package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/wangle201210/studyGo/tools/bilibili/lib"
	"github.com/wangle201210/studyGo/tools/bilibili/response"
)

const (
	UrlUpstat = BaseUrl + "/x/space/upstat?mid=%d&jsonp=jsonp"
)

func (s *Service) UpdateUpstat(ctx context.Context) (err error) {
	u := s.dal.User
	users, err := u.WithContext(ctx).Where(u.ID.Gte(106)).Find()
	if err != nil {
		log.Errorc(ctx, "Find err %+v", err)
		return
	}
	for _, user := range users {
		if err = s.GetUpstat(ctx, user.Mid); err != nil {
			log.Errorc(ctx, "GetUpstat err %+v", err)
			return
		}
	}
	return
}
func (s *Service) GetUpstat(ctx context.Context, mid int32) (err error) {
	url := fmt.Sprintf(UrlUpstat, mid)
	res := new(response.Upstat)
	if err = lib.HttpGet(url, &lib.HttpParam{Cookie: userCookie}, &res); err != nil {
		log.Errorc(ctx, "HttpGet err %+v", err)
		return
	}
	log.Infoc(ctx, "GetUserByName res (%+v)", res)
	data := res.Data
	i := s.dal.Info
	_, err = i.WithContext(ctx).Where(i.Mid.Eq(mid)).
		Assign(i.ArchiveView.Value(int64(data.Archive.View)), i.Likes.Value(int64(data.Likes))).FirstOrCreate()
	if err != nil {
		log.Errorc(ctx, "FirstOrCreate err %+v", err)
		return
	}
	return
}
