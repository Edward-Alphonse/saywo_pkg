package http

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type HttpGetLoader[T any] struct {
	ctx     context.Context
	getUrl  string
	params  url.Values
	timeout time.Duration

	data *T
}

func NewHttpGetLoader[T any](ctx context.Context, url string, params url.Values, timeout time.Duration) *HttpGetLoader[T] {
	return &HttpGetLoader[T]{
		ctx:     ctx,
		getUrl:  url,
		params:  params,
		timeout: timeout,
	}
}

func (l *HttpGetLoader[T]) Load() error {
	url, err := url.Parse(l.getUrl)
	if err != nil {
		return err
	}
	url.RawQuery = l.params.Encode()

	client := &http.Client{
		Timeout: l.timeout,
	}
	resp, err := client.Get(url.String())

	if err != nil {
		return errors.Wrap(err, "compareFacehandler compare failed")
	}
	defer resp.Body.Close()

	contents, _ := ioutil.ReadAll(resp.Body)
	result := new(T)
	err = json.Unmarshal(contents, result)
	if err != nil {
		return errors.Wrap(err, "CompareFaceLoader load failed")
	}
	l.data = result
	return nil
}

func (l *HttpGetLoader[T]) GetData() *T {
	return l.data
}
