package service

import (
	"github.com/pkg/errors"
)

func (s *Service) CheckTask(id string) (AssignTaskResp, error) {
	if s.Cache == nil {
		return AssignTaskResp{
			Id:     id,
			Status: statusError,
		}, errors.New(cacheNotInitializedErrMsg)
	}
	
	item, found := s.Cache.Get(id)
	if !found {
		return AssignTaskResp{
			Id:     id,
			Status: statusError,
		}, errors.New(notFoundErrMsg)
	}
	
	switch item.(type) {
	case AssignTaskResp:
		return item.(AssignTaskResp), nil
	case AssignTaskReq:
		return AssignTaskResp{}, nil
	default:
		return AssignTaskResp{}, errors.New("unknown type")
	}
}

