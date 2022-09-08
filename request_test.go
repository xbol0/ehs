package ehs

import (
	"testing"
	"time"
)

func TestRequest(t *testing.T) {
	result, err := Request(
		&RequestConfig{
			Host:    "www.ithome.com",
			Url:     "https://182.106.137.35/",
			Timeout: time.Second * 2,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", result)
}

func TestRequestSkipFail(t *testing.T) {
	result, err := Request(
		&RequestConfig{
			Host:     "xxx.ithome.com",
			Url:      "https://182.106.137.35/",
			Timeout:  time.Second * 2,
			SkipFail: true,
		},
	)

	if err != ErrSkipFail {
		t.Fatal(err)
	}

	t.Logf("%v", result)
}
