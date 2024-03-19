package http

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpPostLoader[T any] struct {
	ctx         context.Context
	params      map[string]any
	contentType string
	postUrl     string
	timeout     time.Duration

	data *T
}

func NewHttpPostLoader[T any](ctx context.Context, postUrl string, params map[string]any, contentType string, timeout time.Duration) *HttpPostLoader[T] {
	return &HttpPostLoader[T]{
		ctx:         ctx,
		params:      params,
		contentType: contentType,
		postUrl:     postUrl,
		data:        new(T),
		timeout:     timeout,
	}
}

func (l *HttpPostLoader[T]) Load() error {
	bytesData, _ := json.Marshal(l.params)
	client := &http.Client{
		Timeout: l.timeout,
	}
	resp, err := client.Post(l.postUrl, l.contentType, bytes.NewBuffer((bytesData)))
	if err != nil {
		return errors.Wrap(err, "HttpPostLoader load failed")
	}

	defer resp.Body.Close()
	result := new(T)
	contents, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(contents, result)
	if err != nil {
		return err
	}
	l.data = result
	return nil
}

func (l *HttpPostLoader[T]) GetData() *T {
	return l.data
}
