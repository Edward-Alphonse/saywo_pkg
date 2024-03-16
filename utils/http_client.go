package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/Edward-Alphonse/saywo_pkg/logs"
)

// Get 发起一个GET请求
func Get(url string, headers map[string]string) (response []byte, err error) {
	client := http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Post 发起一个Post请求
func Post(url string, data interface{}, contentType string) (content []byte, err error) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", contentType)
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Put 发起一个Put请求
func Put(url string, data interface{}, contentType string) ([]byte, error) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", contentType)
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostJSON 以json模式发起一个POST请求
func PostJSON(header map[string]string, body map[string]interface{}, url string) (map[string]interface{}, error) {
	httpBodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(httpBodyBytes))
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("post request failed")
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	logs.InfoKv("", "", "PostJson response body,[data=%s]", string(respBody))

	respBodyMap := make(map[string]interface{})
	err = json.Unmarshal(respBody, &respBodyMap)
	if err != nil {
		return nil, err
	}
	return respBodyMap, nil
}

// PostAny 以json模式发起一个POST请求
func PostAny(header map[string]string, any interface{}, url string) ([]byte, error) {
	logs.InfoKv("", "", "PostAny request url,[url=%s]", url)
	httpBodyBytes, err := json.Marshal(any)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(httpBodyBytes))
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		logs.InfoKv("", "", "PostAny header,[key=%s,value=%s]", key, value)
		req.Header.Set(key, value)
	}
	logs.InfoKv("", "", "PostAny request body,[data=%s]", string(httpBodyBytes))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	logs.InfoKv("", "", "PostAny response status,[status=%s]", resp.Status)
	if resp.StatusCode != 200 {
		respBody, _ := io.ReadAll(resp.Body)
		logs.ErrorByArgs("post request failed,body:%s", string(respBody))
		return nil, errors.New("post request failed")
	}

	respBody, err := io.ReadAll(resp.Body)
	logs.InfoKv("", "", "PostAny response body,[data=%s]", string(respBody))
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// PostForm 以form模式发起一个POST请求
func PostForm(header map[string]string, formData url.Values, url string) ([]byte, error) {
	logs.InfoKv("", "", "PostAny request url,[url=%s]", url)
	resp, err := http.PostForm(url, formData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	logs.InfoKv("", "", "PostAny response status,[status=%s]", resp.Status)
	if resp.StatusCode != 200 {
		respBody, _ := io.ReadAll(resp.Body)
		logs.ErrorByArgs("post request failed,body:%s", string(respBody))
		return nil, errors.New("post request failed")
	}

	respBody, err := io.ReadAll(resp.Body)
	logs.InfoKv("", "", "PostAny response body,[data=%s]", string(respBody))
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// PostFormWith201 以form模式发起一个POST请求
func PostFormWith201(formData url.Values, url string) ([]byte, error) {
	logs.InfoKv("", "", "PostAny request url,[url=%s]", url)
	resp, err := http.PostForm(url, formData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	logs.InfoKv("", "", "PostAny response status,[status=%s]", resp.Status)
	if resp.StatusCode != 201 {
		respBody, _ := io.ReadAll(resp.Body)
		logs.ErrorByArgs("post request failed,body:%s", string(respBody))
		return nil, errors.New("post request failed")
	}

	respBody, err := io.ReadAll(resp.Body)
	logs.InfoKv("", "", "PostAny response body,[data=%s]", string(respBody))
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// PostFormData 以form模式发起一个POST请求  form-data
func PostFormData(header map[string]string, values map[string]string, url string) ([]byte, error) {
	logs.InfoKv("", "", "PostAny request url,[url=%s]", url)
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range values {
		_ = w.WriteField(k, v)
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		logs.InfoKv("", "", "PostAny header,[key=%s,value=%s]", key, value)
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	logs.InfoKv("", "", "PostAny response status,[status=%s]", resp.Status)
	if resp.StatusCode != 200 {
		respBody, _ := io.ReadAll(resp.Body)
		logs.ErrorByArgs("post request failed,body:%s", string(respBody))
		return nil, errors.New("post request failed")
	}

	respBody, err := io.ReadAll(resp.Body)
	logs.InfoKv("", "", "PostAny response body,[data=%s]", string(respBody))
	if err != nil {
		return nil, err
	}
	logs.InfoKv("", "", "PostJson response body,[data=%s]", string(respBody))

	return respBody, nil
}
