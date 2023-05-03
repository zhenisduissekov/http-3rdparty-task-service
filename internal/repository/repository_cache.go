package repository

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/zhenisduissekov/http-3rdparty-task-service/config"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/entity"
)

const (
	failedToCloseRespBody = "failed to close response body"
	noTaskFoundErrMsg     = "no task found"
	uknownTypeErrMsg      = "unknown type error"
)

type RepoCache struct {
	cache      *cache.Cache
	cnf        *config.Conf
	httpClient *http.Client
}

func NewCache(cnf *config.Conf) *RepoCache {
	newCache := cache.New(cnf.Cache.DefaultExpiration, cnf.Cache.CleanupInterval)
	return &RepoCache{
		cnf:   cnf,
		cache: newCache,
		httpClient: &http.Client{
			Timeout: cnf.Auth.RequestTimeout,
		},
	}
}

func (r *RepoCache) Set(key string, value entity.Task) {
	r.cache.Set(key, value, r.cnf.Cache.DefaultExpiration)
}

func (r *RepoCache) Get(key string) (entity.Task, error) {
	val, ok := r.cache.Get(key)
	if !ok {
		return entity.Task{}, errors.New(noTaskFoundErrMsg)
	}
	task, ok := val.(entity.Task)
	if !ok {
		return entity.Task{}, errors.New(uknownTypeErrMsg)
	}
	return task, nil
}

func (r *RepoCache) MakeRequest(items entity.Task) ([]byte, map[string]string, int, error) {
	req, err := prepReq(items)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	resp, err := r.httpClient.Do(req)
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
