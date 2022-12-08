package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/wangle201210/studyGo/tools/bilibili/lib"
	"github.com/wangle201210/studyGo/tools/bilibili/model"
	"github.com/wangle201210/studyGo/tools/bilibili/response"
	"net/http"
	"time"
)

const (
	BaseUrl       = "https://api.bilibili.com"
	UrlUser       = BaseUrl + "/x/space/acc/info?mid=%d&jsonp=jsonp"
	UrlUserSearch = BaseUrl + "/x/web-interface/search/type?search_type=bili_user&keyword=%s"
)

var (
	userCookie = &http.Cookie{
		Name:   "SESSDATA",
		Value:  "102a7464,1680189056,cdb19*a1",
		Domain: BaseUrl,
	}
	searchCookie = &http.Cookie{
		Name:   "buvid3",
		Value:  "14CDFDBB-2233-45C2-8472-35E7F5C0E2D118567infoc",
		Domain: BaseUrl,
	}
)

func (s *Service) CheckIfExist(ctx context.Context, name string) (exist bool, err error) {
	u := s.dal.User
	c, err := s.dal.User.WithContext(ctx).
		Where(u.Name.Eq(name)).Count()
	if err != nil {
		log.Errorc(ctx, "Count err %+v", err)
		return
	}
	exist = c > 0
	return
}

func (s *Service) UpdateUerList(ctx context.Context, list []string) (err error) {
	var noData []string
	var exist bool
	for _, name := range list {
		exist, err = s.CheckIfExist(ctx, name)
		if err != nil {
			log.Errorc(ctx, "CheckIfExist err %+v", err)
			return
		}
		if exist {
			continue
		}
		var mid int
		mid, err = s.GetUserMidByName(ctx, name)
		if err != nil {
			log.Errorc(ctx, "GetUserMidByName %+v", err)
			return
		}
		if mid == 0 {
			log.Errorc(ctx, "GetUserMidByName name (%s) not found mid", name)
			noData = append(noData, name)
			continue
		}
		err = s.SaveUserInfo(ctx, int32(mid))
		if err != nil {
			log.Errorc(ctx, "SaveUserInfo %+v", err)
			return
		}
		// 不慢点，会被限制访问
		time.Sleep(time.Second)
	}
	if len(noData) > 0 {
		log.Infoc(ctx, "UpdateUerList noData list %+v", noData)
	}
	return
}

func (s *Service) GetUserMidByName(ctx context.Context, name string) (mid int, err error) {
	url := fmt.Sprintf(UrlUserSearch, name)
	res := new(response.NameSearch)
	if err = lib.HttpGet(url, &lib.HttpParam{Cookie: searchCookie}, &res); err != nil {
		log.Errorc(ctx, "HttpGet err %+v", err)
		return
	}
	log.Infoc(ctx, "GetUserByName res (%+v)", res)
	for _, item := range res.Data.Result {
		if item.Uname == name {
			mid = item.Mid
			return
		}
	}
	return
}

func (s *Service) SaveUserInfo(ctx context.Context, userId int32) (err error) {
	url := fmt.Sprintf(UrlUser, userId)
	res := new(response.User)
	if err = lib.HttpGet(url, &lib.HttpParam{Cookie: userCookie}, res); err != nil {
		log.Errorc(ctx, "HttpGet err %+v", err)
		return
	}
	log.Infoc(ctx, "GetUserInfo res (%+v)", res)
	data := res.Data
	u := s.dal.User
	user, err := s.dal.User.WithContext(ctx).
		Where(u.Mid.Eq(userId)).
		Assign(
			u.Name.Value(data.Name),
			u.Face.Value(data.Face),
			u.Sex.Value(data.Sex),
		).FirstOrCreate()
	if err != nil {
		log.Errorc(ctx, "FirstOrCreate err %+v", err)
		return
	}
	log.Infoc(ctx, "GetUserInfo user (%+v)", user)
	return
}

func (s *Service) GetUserListInfo(ctx context.Context, offset, limit int) (res []*response.UserInfo, err error) {
	i := s.dal.Info
	page, _, err := i.WithContext(ctx).Order(i.ArchiveView).FindByPage(offset, limit)

	if err != nil {
		log.Errorc(ctx, "FindByPage err %+v", err)
		return
	}
	var mids []int32
	for _, i := range page {
		mids = append(mids, i.Mid)
	}
	u := s.dal.User
	users, err := u.WithContext(ctx).Where(u.Mid.In(mids...)).Find()
	if err != nil {
		log.Errorc(ctx, "Find err %+v", err)
		return
	}
	userMap := make(map[int32]*model.User)
	for _, u := range users {
		userMap[u.Mid] = u
	}
	for _, p := range page {
		i, ok := userMap[p.Mid]
		if !ok {
			continue
		}
		item := &response.UserInfo{
			Mid:     i.Mid,
			Name:    i.Name,
			Sex:     i.Sex,
			Face:    i.Face,
			Archive: p.ArchiveView,
			Likes:   p.Likes,
		}
		res = append(res, item)
	}
	return
}
