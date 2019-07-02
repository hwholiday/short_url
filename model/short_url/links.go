package short_url

import (
	"errors"
	"short_url/base"
	"short_url/base/tool"

	"go.uber.org/zap"
)

type Links struct {
	Id         int    `json:"id"`
	Url        string `json:"url"`
	Keyword    string `json:"keyword"`
	Status     int    `json:"status"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

//生成短链接地址
//status 1系统分配 2用户自定义
func (s *service) CreateLinks(url, key string, status int) (string, error) {
	var (
		err     error
		order   int64
		links   *Links
		keyword string
	)
	//先判断存在该短码没
	if links, err = s.GetLinksByUrl(url); err != nil {
		tool.GetLogger().Error("[CreateLinks] GetLinksByUrl", zap.Error(err))
		return keyword, err
	}
	if links != nil { //存在直接返回
		return links.Keyword, nil
	}
	if status == base.LINKS_CREATE_SYSTEM {
		if order, err = s.sequence.GetBorrowOrder(); err != nil {
			tool.GetLogger().Error("[CreateLinks] GetBorrowOrder", zap.Error(err))
			return keyword, err
		}
		keyword = tool.TenToAny(int(order), 62)
	} else if status == base.LINKS_CREATE_CUSTOM {
		//判断该keyword是否存在
		if links, err = s.GetLinksByKeyword(key); err != nil {
			tool.GetLogger().Error("[CreateLinks] GetLinksByUrl", zap.Error(err))
			return keyword, err
		}
		if links != nil { //存在直接返回
			return keyword, errors.New("already exists")
		}
		keyword = key
	} else {
		tool.GetLogger().Warn("[GetLinksByUrl] undefined", zap.Any("status", status))
		return keyword, errors.New("undefined")
	}
	links = new(Links)
	links.Status = status
	links.Keyword = keyword
	links.CreateTime = base.GetTime()
	links.UpdateTime = base.GetTime()
	links.Url = url
	if _, err = s.m.InsertOne(links); err != nil {
		tool.GetLogger().Warn("[GetLinksByUrl] undefined", zap.Any("status", status))
		return keyword, err
	}
	return keyword, nil
}

func (s *service) GetLinksByUrl(url string) (*Links, error) {
	var (
		err   error
		has   bool
		links = &Links{}
	)
	has, err = s.m.Where("url = ?", url).Get(links)
	if err != nil {
		tool.GetLogger().Error("[GetLinksByUrl] Where", zap.Error(err))
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return links, nil
}

func (s *service) GetLinksByKeyword(keyword string) (*Links, error) {
	var (
		err   error
		has   bool
		links = &Links{}
	)
	has, err = s.m.Where("keyword = ?", keyword).Get(links)
	if err != nil {
		tool.GetLogger().Error("[GetLinksByKeyword] Where", zap.Error(err))
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return links, nil
}
