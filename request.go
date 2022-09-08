package ehs

import (
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	"time"
)

type RequestConfig struct {
	Url      string
	Host     string
	Timeout  time.Duration
	SkipFail bool
}

type SiteMeta struct {
	Status int
	Title  string
	Type   string
}

var ErrSkipFail = errors.New("Skip fail response.")

func Request(c *RequestConfig) (*SiteMeta, error) {
	q, err := http.NewRequest("GET", c.Url, nil)

	if err != nil {
		return nil, err
	}

	q.Host = c.Host
	q.Header.Add("user-agent", UA)

	client := makeClient(c.Timeout)
	res, err := client.Do(q)

	if err != nil {
		return nil, err
	}

	if c.SkipFail && !(res.StatusCode >= 200 && res.StatusCode < 300) {
		return nil, ErrSkipFail
	}

	if res.ContentLength > MAX_SIZE {
		return nil, ERR_TOO_LARGE
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	return &SiteMeta{
		Status: res.StatusCode,
		Title:  getTitle(string(body)),
		Type:   res.Header.Get("content-type"),
	}, nil
}

func makeClient(t time.Duration) *http.Client {
	return &http.Client{
		Timeout: t,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},

		// No redirect.
		// Because if redirect from http to https, the tasks have be in the queue.
		// The other cases may not be we needed.
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}
