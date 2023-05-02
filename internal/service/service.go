package service

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/entity"
)

func (s *NewService) AssignTask(items entity.Task) (string, error) {
	items.Status = statusNew
	s.repository.Set(items.Id, items)
	queue <- items
	return items.Id, nil
}

func (s *NewService) TaskQueue(ctx context.Context) error {
	ticker := time.NewTicker(tickPeriod)
	defer ticker.Stop()

	for {
		select {
		case nextTask, ok := <-queue:
			if !ok {
				log.Warn().Msg(channelWasClosedMsg)
				return nil
			}
			log.Info().Msg(taskReceivedMsg)
			s.processNextTask(nextTask)
		case <-ticker.C:
			log.Debug().Msg(tickMsg)
		case <-ctx.Done():
			s.CloseChannel() //todo: unclear
			return nil
		}
	}

	return nil
}

func (s *NewService) CloseChannel() {
	close(queue)
}

func (s *NewService) processNextTask(items entity.Task) {
	status := statusDone
	body, headers, statusCode, err := s.makeRequest(items)
	if err != nil {
		log.Error().Err(err).Msg(failedToMakeRequestErrMsg)
		status = statusError
	}

	s.repository.Set(items.Id, entity.Task{
		Id:             items.Id,
		Url:            items.Url,
		Method:         items.Method,
		Status:         status,
		HttpStatusCode: statusCode,
		ReqBody:        items.ReqBody,
		RespBody:       string(body),
		Length:         len(body),
		Headers:        headers,
	})
}
func (s *NewService) makeRequest(items entity.Task) ([]byte, map[string]string, int, error) {
	req, err := prepReq(items)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, getRespHeaders(resp), 0, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Warn().Msg(failedToCloseRespBody)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, getRespHeaders(resp), resp.StatusCode, err
	}

	return body, getRespHeaders(resp), resp.StatusCode, nil
}

func prepReq(items entity.Task) (*http.Request, error) {
	req, err := http.NewRequest(items.Method, items.Url, bytes.NewBuffer([]byte(items.ReqBody)))
	if err != nil {
		return nil, err
	}

	for k, v := range items.Headers {
		req.Header.Set(k, v)
	}

	return req, nil
}

func getRespHeaders(resp *http.Response) map[string]string {
	if resp == nil {
		return nil
	}
	headers := make(map[string]string)

	for key, values := range resp.Header {
		headers[key] = strings.Join(values, ",")
	}
	return headers
}

func (s *NewService) CheckTask(id string) (entity.Task, error) {
	return s.repository.Get(id)
}
