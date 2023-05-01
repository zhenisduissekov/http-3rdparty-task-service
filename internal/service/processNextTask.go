package service

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
)

func (s *Service) processNextTask(id string) {
	items, err := s.getItem(id)
	if err != nil {
		s.log.Error().Err(err).Msg(failedToGetItemErrMsg)
		return
	}
	
	req, err := prepReq(items)
	if err!= nil {
		s.log.Error().Err(err).Msg(failedToPrepareReq)
		return
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		s.log.Error().Err(err).Msg(failedRequest)
		return
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			s.log.Warn().Msg(failedToCloseRespBody)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.log.Error().Err(err).Msg(failedToReadRespBody)
		return
	}
	
	response := AssignTaskResp{
		Id:    id,
		Body:  string(body),
		Status: resp.Status,
		Length: len(body),
		Headers: getRespHeaders(resp),
	}

	s.Cache.Set(id, response, cache.DefaultExpiration)

}


func prepReq(items AssignTaskReq) (*http.Request, error) {
	req, err := http.NewRequest(items.Method, items.Url, bytes.NewBuffer(items.ReqBody))
	if err != nil {
		return nil, err
	}
	
	for k, v := range items.Headers {
		req.Header.Set(k, v)
	}

	return req, nil
}

func getRespHeaders(resp *http.Response) map[string]string {
	headers := make(map[string]string)
    for key, values := range resp.Header {
        headers[key] = strings.Join(values, ",")
    }
	return headers
}


func (s *Service) getItem(id string) (AssignTaskReq, error) {
	item, ok := s.Cache.Get(id)
	if !ok {
		return AssignTaskReq{}, errors.New(notFoundErrMsg)
	}

	items, ok := item.(AssignTaskReq)
	if !ok {
		return AssignTaskReq{}, errors.New(itemIsNotAssignReq)
	}
	
	return items, nil
}